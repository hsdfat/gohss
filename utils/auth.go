/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"gohss/messages"

	// "gohss/message"
	// "gohss/models"

	"github.com/emakeev/milenage"
	"github.com/golang/glog"
	"github.com/omec-project/openapi/models"
)

const (
	// indBits is the number of bits reserved for IND (one of the two parts of SQN).
	// See 3GPP TS 33.102 Appendix C.1.1.1 and C.3.
	indBits = 5

	// indMask is a bit mask where a bit is 1 if and only if it is a part of ind.
	indMask = (1 << indBits) - 1

	// seqMask is a bit mask where a bit is 1 if and only if it is a part of seq.
	seqMask = (1 << 48) - 1 - indMask

	// lteResyncInfoBytes is the expected size of the lte resync info in bytes.
	// The first 16 bytes store RAND and the next 14 bytes store AUTS.
	lteResyncInfoBytes = milenage.RandChallengeBytes + milenage.ExpectedAutsBytes

	// maxSeqDelta is the maximum allowed increase to SQN.
	// eg. if x was the last accepted SQN, then the next SQN must
	// be greater than x and less than (x + maxSeqDelta) to be accepted.
	// See 3GPP TS 33.102 Appendix C.2.1.
	maxSeqDelta = 1 << 28

	maxReturnedVectors = 5
)

// GenerateLteAuthVectors generates at most `numVectors` lte auth vectors.
// Inputs:
//   - numVectors  -- The maximum number of vectors to generate
//   - milenage    -- The cipher to use to generate the vector
//   - subscriber  -- The subscriber data for the subscriber we want to generate auth vectors for
//   - plmn        -- 24 bit network identifier
//   - authSqnInd  -- the IND of the current vector being generated
//
// Returns:
//   - The E-UTRAN vectors, UTRAN vectors, the next value to set the subscriber's LteAuthNextSeq to (or an error)
func GenerateLteAuthVectors(
	numEutranVectors uint32,
	numUtranVectors uint32,
	mcipher *milenage.Cipher,
	subscriber *models.AuthenticationSubscription,
	plmn,
	lteAuthOp []byte,
	authSqnInd uint64) ([]*milenage.EutranVector, []*milenage.UtranVector, uint64, error) {

	if numEutranVectors > maxReturnedVectors {
		numEutranVectors = maxReturnedVectors
	}
	var vectors = make([]*milenage.EutranVector, 0, numEutranVectors)
	lteAuthNextSeq, err := hexToInt(subscriber.SequenceNumber, 12)
	if err != nil {
		return nil, nil, 0, err
	}
	for i := uint32(0); i < numEutranVectors; i++ {
		vector, nextSeq, err := GenerateLteAuthVector(mcipher, subscriber, plmn, lteAuthOp, authSqnInd)
		if err != nil {
			// If we have already generated an auth vector successfully, then we can
			// return it. Otherwise, we must signal an error.
			// See 3GPP TS 29.272 section 5.2.3.1.3.
			if i == 0 {
				return nil, nil, 0, err
			}
			glog.Errorf("failed to generate lte auth vector: %v", err)
			break
		}
		vectors = append(vectors, vector)
		lteAuthNextSeq = nextSeq
		subscriber.SequenceNumber = IntToHexFixedLength(lteAuthNextSeq, 12)
	}
	if numUtranVectors > maxReturnedVectors {
		numUtranVectors = maxReturnedVectors
	}
	var utranVectors = make([]*milenage.UtranVector, 0, numUtranVectors)
	for i := uint32(0); i < numUtranVectors; i++ {
		vector, nextSeq, err := GenerateUtranAuthVector(mcipher, subscriber, lteAuthOp, authSqnInd)
		if err != nil {
			// If we have already generated an auth vector successfully, then we can
			// return it. Otherwise, we must signal an error.
			// See 3GPP TS 29.272 section 5.2.3.1.3.
			if len(vectors) == 0 && i == 0 {
				return nil, nil, 0, err
			}
			glog.Errorf("failed to generate UTRAN auth vector: %v", err)
			break
		}
		lteAuthNextSeq = nextSeq
		subscriber.SequenceNumber = IntToHexFixedLength(lteAuthNextSeq, 12)
		utranVectors = append(utranVectors, vector)
	}

	return vectors, utranVectors, lteAuthNextSeq, nil
}

// GenerateLteAuthVector returns the lte auth vector for the subscriber.
// Inputs:
//   - milenage    -- The cipher to use to generate the vector
//   - subscriber  -- The subscriber data for the subscriber we want to generate auth vectors for
//   - plmn        -- 24 bit network identifier
//   - authSqnInd  -- the IND of the current vector being generated
//
// Returns:
//   - A E-UTRAN vector and the next value to set the subscriber's LteAuthNextSeq to (or an error).
func GenerateLteAuthVector(
	mcipher *milenage.Cipher,
	subs *models.AuthenticationSubscription,
	plmn, lteAuthOp []byte,
	authSqnInd uint64) (*milenage.EutranVector, uint64, error) {

	// if err := ValidateLteSubscription(subs); err != nil {
	// 	return nil, 0, messages.NewAuthRejectedError(err.Error())
	// }
	// if subs.State == nil {
	// 	return nil, 0, message.NewAuthRejectedError("Subscriber data missing subscriber state")
	// }
	sqn, err := hexToInt(subs.SequenceNumber, 12)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}

	key, err := hex.DecodeString(subs.PermanentKey.PermanentKeyValue)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}

	opc, err := GetOrGenerateOpc(subs, key, lteAuthOp)
	if err != nil {
		return nil, 0, err
	}

	vector, err := mcipher.GenerateEutranVector(key, opc, sqn, plmn)
	if err != nil {
		return vector, 0, messages.NewAuthRejectedError(err.Error())
	}
	return vector, sqn + 1, err
}

// GenerateSIPAuthVectors generates at most `numVectors` lte auth vectors.
// Inputs:
//   - numVectors  -- The maximum number of vectors to generate
//   - milenage    -- The cipher to use to generate the vector
//   - subscriber  -- The subscriber data for the subscriber we want to generate auth vectors for
//   - authSqnInd  -- the IND of the current vector being generated
//
// Returns:
//   - The E-UTRAN vectors,  the next value to set the subscriber's SIPAuthNextSeq to (or an error)
func GenerateSIPAuthVectors(
	numSIPVectors uint32,
	mcipher *milenage.Cipher,
	subscriber *models.AuthenticationSubscription,
	lteAuthOp []byte,
	authSqnInd uint64) ([]*milenage.SIPAuthVector, uint64, error) {

	if numSIPVectors > maxReturnedVectors {
		numSIPVectors = maxReturnedVectors
	}
	var vectors = make([]*milenage.SIPAuthVector, 0, numSIPVectors)
	lteAuthNextSeq, err := hexToInt(subscriber.SequenceNumber, 12)
	if err != nil {
		return nil, 0, err
	}
	for i := uint32(0); i < numSIPVectors; i++ {
		vector, nextSeq, err := GenerateSIPAuthVector(mcipher, subscriber, lteAuthOp, authSqnInd)
		if err != nil {
			// If we have already generated an auth vector successfully, then we can
			// return it. Otherwise, we must signal an error.
			// See 3GPP TS 29.272 section 5.2.3.1.3.
			if i == 0 {
				return nil, 0, err
			}
			glog.Errorf("failed to generate lte auth vector: %v", err)
			break
		}
		vectors = append(vectors, vector)
		lteAuthNextSeq = nextSeq
		subscriber.SequenceNumber = IntToHexFixedLength(lteAuthNextSeq, 12)
	}
	

	return vectors, lteAuthNextSeq, nil
}

// GenerateSIPAuthVector returns the sip auth vector for the subscriber.
// Inputs:
//   - milenage    -- The cipher to use to generate the vector
//   - subscriber  -- The subscriber data for the subscriber we want to generate auth vectors for
//   - authSqnInd  -- the IND of the current vector being generated
//
// Returns:
//   - A SIP vector and the next value to set the subscriber's LteAuthNextSeq to (or an error).
func GenerateSIPAuthVector(
	mcipher *milenage.Cipher,
	subs *models.AuthenticationSubscription,
	lteAuthOp []byte,
	authSqnInd uint64) (*milenage.SIPAuthVector, uint64, error) {

	// if err := ValidateLteSubscription(subs); err != nil {
	// 	return nil, 0, messages.NewAuthRejectedError(err.Error())
	// }
	// if subs.State == nil {
	// 	return nil, 0, message.NewAuthRejectedError("Subscriber data missing subscriber state")
	// }
	sqn, err := hexToInt(subs.SequenceNumber, 12)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}

	key, err := hex.DecodeString(subs.PermanentKey.PermanentKeyValue)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}

	opc, err := GetOrGenerateOpc(subs, key, lteAuthOp)
	if err != nil {
		return nil, 0, err
	}

	vector, err := mcipher.GenerateSIPAuthVector(key, opc, sqn)
	if err != nil {
		return vector, 0, messages.NewAuthRejectedError(err.Error())
	}
	return vector, sqn + 1, err
}

// GenerateUtranAuthVector returns the lte auth vector for the subscriber.
// Inputs:
//   - milenage    -- The cipher to use to generate the vector
//   - subscriber  -- The subscriber data for the subscriber we want to generate auth vectors for
//   - plmn        -- 24 bit network identifier
//   - authSqnInd  -- the IND of the current vector being generated
//
// Returns:
//   - A UTRAN vector and the next value to set the subscriber's LteAuthNextSeq to (or an error).
func GenerateUtranAuthVector(
	mcipher *milenage.Cipher,
	subs *models.AuthenticationSubscription,
	lteAuthOp []byte,
	authSqnInd uint64) (*milenage.UtranVector, uint64, error) {

	// lte := subscriber.Lte
	// if err := ValidateLteSubscription(subs); err != nil {
	// 	return nil, 0, messages.NewAuthRejectedError(err.Error())
	// }
	key, err := hex.DecodeString(subs.PermanentKey.PermanentKeyValue)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}
	opc, err := GetOrGenerateOpc(subs, key, lteAuthOp)
	if err != nil {
		return nil, 0, err
	}
	sqn, err := hexToInt(subs.SequenceNumber, 12)
	if err != nil {
		return nil, 0, messages.NewAuthRejectedError(err.Error())
	}
	sqn_real := SeqToSqn(sqn, authSqnInd)
	vector, err := mcipher.GenerateUtranVector(key, opc, sqn_real)
	if err != nil {
		return vector, 0, messages.NewAuthRejectedError(err.Error())
	}
	return vector, sqn + 1, err
}

// ResyncLteAuthSeq validates a re-synchronization request and computes the SEQ
// from the AUTS sent by U-SIM. The next value of lteAuthNextSeq (or an error) is returned.
// See 3GPP TS 33.102 section 6.3.5.
func ResyncLteAuthSeq(subs *models.AuthenticationSubscription, resyncInfo, lteAuthOp []byte) (uint64, error) {

	// if IsAllZero(resyncInfo) {
	// 	return subs.State.SQN, nil
	// }
	if len(resyncInfo) != lteResyncInfoBytes {
		err := messages.NewAuthRejectedError(fmt.Sprintf("resync info incorrect length. expected %v bytes, but got %v bytes", lteResyncInfoBytes, len(resyncInfo)))
		return 0, err
	}

	// if err := ValidateLteSubscription(subs); err != nil {
	// 	return 0, messages.NewAuthRejectedError(err.Error())
	// }

	// Use dummy AMF for re-synchronization. See 3GPP TS 33.102 section 6.3.3.
	mcipher, err := milenage.NewCipher(make([]byte, milenage.ExpectedAmfBytes))
	if err != nil {
		return 0, messages.NewAuthDataUnavailableError(err.Error())
	}
	rand := resyncInfo[:milenage.RandChallengeBytes]
	auts := resyncInfo[milenage.RandChallengeBytes:]
	key, err := hex.DecodeString(subs.PermanentKey.PermanentKeyValue)
	if err != nil {
		return 0, messages.NewAuthRejectedError(err.Error())
	}
	opc, err := GetOrGenerateOpc(subs, key, lteAuthOp)
	if err != nil {
		return 0, err
	}
	sqnMs, macS, err := mcipher.GenerateResync(auts, key, opc, rand)
	if err != nil {
		return 0, messages.NewAuthDataUnavailableError(err.Error())
	}
	if !bytes.Equal(macS[:], auts[milenage.ExpectedAutsBytes-len(macS):]) {
		return 0, messages.NewAuthRejectedError("Invalid resync authentication code")
	}

	return GetNextLteAuthSqnAfterResync(subs, sqnMs)
}

// GetNextLteAuthSqnAfterResync returns the value of the next sequence number after
// sqn or an error if a resync should not occur.
// See 3GPP TS 33.102 Appendix C.3.
func GetNextLteAuthSqnAfterResync(subs *models.AuthenticationSubscription, sqn uint64) (uint64, error) {
	// if state == nil {
	// 	return 0, messages.NewAuthDataUnavailableError("subscriber state was nil")
	// }
	sqnOld, _ := hexToInt(subs.SequenceNumber, 12)

	seq, _ := SplitSqn(sqn)
	currentSeq := sqnOld - 1
	if seq < currentSeq {
		seqDelta := currentSeq - seq
		if seqDelta <= maxSeqDelta {
			// This error indicates that the last sequence number should have been
			// accepted by the USIM but wasn't (this should never happen).
			return 0, messages.NewAuthRejectedError(fmt.Sprintf("Re-sync delta in range but UE rejected auth: %d", seqDelta))
		}
	}

	return seq + 1, nil
}

// ValidateLteSubscription returns an error if and only if the lte proto is not
// configured up to use the milenage authentication algorithm.
func ValidateLteSubscription(subs *models.AuthenticationSubscription) error {
	if subs == nil {
		return fmt.Errorf("Subscriber data missing LTE subscription")
	}
	// if lte.State != protos.LTESubscription_ACTIVE {
	// 	return fmt.Errorf("LTE Service not active")
	// }
	// if subs.AuthAlgo != models.MILENAGE {
	// 	return fmt.Errorf("Unsupported milenage algorithm: %v", subs.AuthAlgo)
	// }
	return nil
}

// GetOrGenerateOpc returns lte.AuthOpc and generates if it isn't stored in the proto
func GetOrGenerateOpc(subs *models.AuthenticationSubscription, key []byte, lteAuthOp []byte) ([]byte, error) {
	if subs == nil || len(subs.Opc.OpcValue) == 0 {
		opc, err := milenage.GenerateOpc(key, lteAuthOp)
		if err != nil {
			err = messages.NewAuthDataUnavailableError(err.Error())
		}
		return opc[:], err
	}
	opc, err := hex.DecodeString(subs.Opc.OpcValue)
	if err != nil {
		return nil, messages.NewAuthDataUnavailableError(err.Error())
	}
	return opc, nil
}

// SeqToSqn computes the 48 bit SQN given a seq given the formula defined in
// 3GPP TS 33.102 Annex C.3.2. The length of IND is 5 bits.
// SQN = SEQ || IND
// Inputs:
//   - seq    -- the sequence number
//   - index  -- the index of the current vector being generated
//
// Output:
//   - The 48 bit SQN
func SeqToSqn(seq, index uint64) uint64 {
	return (seq << indBits & seqMask) + (index & indMask)
}

// SplitSqn computes the SEQ and IND given a 48 bit SQN using the formula defined in
// 3GPP TS 33.102 Annex C.3.2. The length of IND is 5 bits.
// SQN = SEQ || IND
// Inputs:
//   - seq   -- the 48 bit SQN
//
// Outputs:
//   - -SEQ and IND
func SplitSqn(sqn uint64) (uint64, uint64) {
	return sqn >> indBits, sqn & indMask
}

// IsAllZero returns true if and only if the slice contains only zero bytes.
func IsAllZero(bytes []byte) bool {
	for _, b := range bytes {
		if b != 0 {
			return false
		}
	}
	return true
}
