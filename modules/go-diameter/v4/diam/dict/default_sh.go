// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package dict

var shXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
    <!--
        3GPP TS 29.329
        https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf
    -->
    <application id="16777217" type="auth" name="SH">
        <vendor id="10415" name="TGPP"/>
        <command code="306" short="UD" name="User-Data">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Server-Name" required="false" max="1"/>
                <rule avp="User-Identity" required="true" max="1"/>

                <rule avp="Service-Indication" required="false"/>
                <rule avp="Data-Reference" required="true"/>
                <rule avp="Identity-Set" required="false" />
                <rule avp="Requested-Domain" required="false" max="1"/>
                <rule avp="Current-Location" required="false" max="1"/>
                <rule avp="DSAI-Tag" required="false"/>
                <rule avp="Session-Priority" required="false" max="1"/>
                <rule avp="User-Name" required="false" max="1"/>
                <rule avp="Requested-Nodes" required="false" max="1"/>
                <rule avp="Serving-Node-Indication" required="false" max="1"/>
                <rule avp="Pre-paging-Supported" required="false" max="1"/>
                <rule avp="Local-Time-Zone-Indication" required="false" max="1"/>
                <rule avp="UDR-Flags" required="false" max="1"/>
                <rule avp="Call-Reference-Info" required="false" max="1"/>

                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Data" required="false" max="1"/>
            </answer>
        </command>
        <command code="307" short="PU" name="Profile-Update">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Data-Reference" required="false"/>
                <rule avp="User-Identity" required="true" max="1"/>
                <rule avp="User-Name" required="false" max="1"/>
                <rule avp="User-Data" required="true" max="1"/>


                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Data" required="false" max="1"/>
                <rule avp="Data-Reference" required="false" max="1"/>
                <rule avp="Repository-Data-ID" required="false" max="1"/>
            </answer>
        </command>
        <command code="308" short="SN" name="Subscribe-Notifications">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Identity" required="true" max="1"/>
                <rule avp="Server-Name" required="false" max="1"/>

                <rule avp="Service-Indication" required="false"/>
                <rule avp="Data-Reference" required="true"/>
                <rule avp="Identity-Set" required="false" />
                <rule avp="DSAI-Tag" required="false"/>
                <rule avp="Session-Priority" required="false" max="1"/>
                <rule avp="One-Time-Notification" required="false" max="1"/>
                <rule avp="Subs-Req-Type" required="true" max="1"/>
                <rule avp="Send-Data-Indication" required="false" max="1"/>
                <rule avp="Expiry-Time" required="false" max="1"/>

                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Data" required="false" max="1"/>
                <rule avp="Expiry-Time" required="false" max="1"/>
            </answer>
        </command>
        <command code="309" short="PN" name="Push-Notification">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Identity" required="true" max="1"/>
                <rule avp="User-Data" required="true" max="1"/>
                <rule avp="User-Name" required="false" max="1"/>

                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129300_129399/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Data" required="true" max="1"/>
                <rule avp="User-Name" required="false" max="1"/>            </answer>
        </command>
       
        <avp name="User-Name" code="1" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

        <avp name="Visited-Network-Identifier" code="600" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="User-Authorization-Type" code="623"  must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>

		<avp name="Server-Capabilities" code="603" must="M,V" may="P" may-encrypt="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Mandatory-Capability" required="false" />
				<rule avp="Optional-Capability" required="false" />
				<rule avp="Server-Name" required="false" />
			</data>
		</avp>


        <avp name="User-Identity" code="700" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
				<rule avp="MSISDN" required="false" />
				<rule avp="Public-Identity" required="false" />
				<rule avp="External-Identifier" required="false" />
			</data>
        </avp>
        <avp name="MSISDN" code="701" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="User-Data" code="702" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Data-Reference" code="703" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="RepositoryData"/>
                <item code="10" name="IMSPublicIdentity"/>
                <item code="11" name="IMSUserState"/>
                <item code="12" name="S-CSCFName"/>
                <item code="13" name="InitialFilterCriteria"/>
                <item code="14" name="LocationInformation"/>
                <item code="15" name="UserState"/>
                <item code="16" name="ChargingInformation"/>
                <item code="17" name="MSISDN"/>
                <item code="18" name="PSIActivation"/>
                <item code="19" name="DSAI"/>
                <item code="21" name="ServiceLevelTraceInfo"/>
                <item code="22" name="IPAddressSecureBindingInformation"/>
                <item code="23" name="ServicePriorityLevel"/>
                <item code="24" name="SMSRegistrationInfo"/>
                <item code="25" name="UEReachabilityForIP"/>
                <item code="26" name="TADSinformation"/>
                <item code="27" name="STN-SR"/>
                <item code="28" name="UE-SRVCC-Capability"/>
                <item code="29" name="ExtendedPriority"/>
                <item code="30" name="CSRN"/>
                <item code="31" name="ReferenceLocationInformation"/>
                <item code="32" name="IMSI"/>
                <item code="33" name="IMSPrivateUserIdentity"/>
            </data>
        </avp>
        <avp name="Service-Indication" code="704" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Subs-Req-Type" code="705" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="Subscribe"/>
                <item code="1" name="Unsubscribe"/>
            </data>
        </avp>

        <avp name="Requested-Domain" code="706" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
           <data type="Enumerated">
                <item code="0" name="CS-Domain"/>
                <item code="1" name="PS-Domain"/>
            </data>
        </avp>
        <avp name="Current-Location" code="707" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="DoNotNeedInitiateActiveLocationRetrieval"/>
                <item code="1" name="InitiateActiveLocationRetrieval"/>
            </data>
        </avp>
        <avp name="Identity-Set" code="708" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="ALL_IDENTITIES"/>
                <item code="1" name="REGISTERED_IDENTITIES"/>
                <item code="2" name="IMPLICIT_IDENTITIES"/>
                <item code="3" name="ALIAS_IDENTITIES"/>

            </data>
        </avp>
        <avp name="Expiry-Time" code="709" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Time"/>
        </avp>
        <avp name="Send-Data-Indication" code="710" must="V"  must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>
        <avp name="DSAI-Tag" code="711" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="One-Time-Notification" code="712" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>
        <avp name="Requested-Nodes" code="713" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Serving-Node-Indication" code="714" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>
        <avp name="Repository-Data-ID" code="715" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
				<rule avp="MSISDN" required="false" />
				<rule avp="Public-Identity" required="false" />
				<rule avp="External-Identifier" required="false" />
			</data>
        </avp>
        <avp name="Sequence-Number" code="716" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Pre-paging-Supported" code="717" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>
        <avp name="Local-Time-Zone-Indication" code="718" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>
            </data>
        </avp>
        <avp name="UDR-Flags" code="719" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Call-Reference-Info" code="720" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
				<rule avp="MSISDN" required="false" />
				<rule avp="Public-Identity" required="false" />
				<rule avp="External-Identifier" required="false" />
			</data>
        </avp>
        <avp name="Call-Reference-Number" code="721" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="AS-Number" code="722" must="V" must-not="M" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Public-Identity" code="601" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="UTF8String"/>
        </avp>

        <avp name="Server-Name" code="602" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.234 Section 10.1.34 -->
            <data type="UTF8String"/>
        </avp>
        <avp name="Supported-Features" code="628" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <!-- 3GPP TS 29.229 Section 6.3.29 -->
            <data type="Grouped">
                <rule avp="Vendor-Id" required="true" max="1"/>
                <rule avp="Feature-List-ID" required="true" max="1"/>
                <rule avp="Feature-List" required="true" max="1"/>
            </data>
        </avp>
        <avp name="Feature-List-ID" code="629" must="V" must_not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
    
        <avp name="Feature-List" code="630" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Session-Priority" code="650" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
        <data type="Enumerated">
            <item code="0" name="PRIORITY-0"/>
            <item code="1" name="PRIORITY-1"/>
            <item code="2" name="PRIORITY-2"/>
            <item code="3" name="PRIORITY-3"/>
            <item code="4" name="PRIORITY-4"/>
        </data>
    </avp>
    </application>
</diameter>`
