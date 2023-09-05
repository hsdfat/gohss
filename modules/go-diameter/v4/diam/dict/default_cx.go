// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package dict

var cxXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
    <!--
        3GPP TS 29.229
        https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf
    -->
    <application id="16777216" type="auth" name="CX">
        <vendor id="10415" name="TGPP"/>
        <command code="300" short="UA" name="User-Authorization">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Public-Identity" required="true" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1" />
                <rule avp="User-Authorization-Type" required="false" max="1" />
                <rule avp="UAR-Flags" required="false"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Server-Name" required="false" max="1"/>
                <rule avp="Server-Capabilities" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="303" short="MA" name="Multimedia-Authentication">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.7 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Server-Name" required="true" max="1"/>
				<rule avp="Public-Identity" required="true" max="1"/>
                <rule avp="RAT-Type" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1" />
                <rule avp="SIP-Auth-Data-Item" required="false" max="1"/>
                <rule avp="SIP-Number-Auth-Items" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
			</request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.7 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="SIP-Number-Auth-Items" required="false" max="1"/>
                <rule avp="SIP-Auth-Data-Item" required="false"/>
                <rule avp="Public-Identity" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="301" short="SA" name="Server-Assignment">
            <request>
            <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Server-Name" required="true" max="1"/>
                <rule avp="User-Data-Already-Available" required="true" max="1"/>
                <rule avp="SCSCF-Restoration-Info" required="false" max="1"/>
                <rule avp="Multiple-Registration-Indication" required="false" max="1"/>
                <rule avp="SAR-Flags" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Server-Assignment-Type" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="User-Data" required="false" max="1"/>
                <rule avp="Charging-Information" required="false" max="1"/>
                <rule avp="Associated-Identities" required="false" max="1"/>
                <rule avp="SCSCF-Restoration-Info" required="false"/>
                <rule avp="Server-Name" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="302" short="LI" name="Location-Info">
            <request>
                <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.1 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Public-Identity" required="true" max="1"/>
                <rule avp="Session-Priority" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1" />
                <rule avp="User-Authorization-Type" required="false" max="1" />
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Server-Name" required="false" max="1"/>
                <rule avp="Server-Capabilities" required="false" max="1"/>
                <rule avp="LIA-Flags" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="304" short="RT" name="Registration-Termination">
            <request>
            <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/17.02.00_60/ts_129229v170200p.pdf Section 6.1.7 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Deregistration-Reason" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.4 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <avp name="Digest-Realm" code="104" must-not="V" must="M" may-encrypt="N">
            <data type="UTF8String" />
        </avp>
        <avp name="Digest-Algorithm" code="111" must-not="V" must="M" may-encrypt="N">
            <data type="UTF8String" />
        </avp>
        <avp name="Digest-QoP" code="110" must-not="V" must="M" may-encrypt="N">
            <data type="UTF8String" />
        </avp>
        <avp name="Digest-HA1" code="121" must-not="V" must="M" may-encrypt="N">
            <data type="UTF8String" />
        </avp>

        <avp name="Service-Selection" code="493" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 5.2.3.5 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Context-Identifier" code="1423" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.27 -->
            <data type="Unsigned32"/>
        </avp>
        <avp name="Session-Id" code="263" must="M" may="P" must-not="V" may-encrypt="Y">
            <data type="UTF8String"/>
        </avp>

        <avp name="Result-Code" code="268" must="M" may="P" must-not="V" may-encrypt="-">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Origin-Host" code="264" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-Realm" code="296" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-State-Id" code="278" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Product-Name" code="269" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Proxy-Host" code="280" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Proxy-Info" code="284" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Proxy-Host" required="true" max="1"/>
				<rule avp="Proxy-State" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Proxy-State" code="33" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="OctetString"/>
		</avp>

        <avp name="Destination-Host" code="293" must="M" may="P" must-not="V" may-encrypt="-">
            <data type="DiameterIdentity"/>
        </avp>

        <avp name="Destination-Realm" code="283" must="M" may="P" must-not="V" may-encrypt="-">
            <data type="DiameterIdentity"/>
        </avp>
		<avp name="Acct-Application-Id" code="259" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Application-Id" code="258" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

        <avp name="Auth-Session-State" code="277" must="M" may="P" must-not="V" may-encrypt="-">
            <data type="Enumerated">
                <item code="0" name="STATE_MAINTAINED"/>
                <item code="1" name="NO_STATE_MAINTAINED"/>
            </data>
        </avp>


		<avp name="User-Name" code="1" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Vendor-Id" code="266" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Specific-Application-Id" code="260" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
			</data>
		</avp>
        <avp name="Visited-Network-Identifier" code="600" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Public-Identity" code="601" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="UTF8String"/>
        </avp>

        <avp name="Server-Name" code="602" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.234 Section 10.1.34 -->
            <data type="UTF8String"/>
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
        <avp name="Mandatory-Capability" code="604" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Optional-Capability" code="605" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="User-Data" code="606" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="SIP-Number-Auth-Items" code="607" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.8 -->
            <data type="Unsigned32" />
        </avp>

        <avp name="SIP-Authentication-Scheme" code="608" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.9 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="SIP-Authenticate" code="609" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.10 -->
            <data type="OctetString"/>
        </avp>

        <avp name="SIP-Authorization" code="610" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.11 -->
            <data type="OctetString"/>
        </avp>
        <avp name="SIP-Authentication-Context" code="611" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.12 -->
            <data type="OctetString"/>
        </avp>
        <avp name="SIP-Auth-Data-Item" code="612" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.13 -->
            <data type="Grouped">
				<rule avp="SIP-Item-Number" required="false" max="1"/>
				<rule avp="SIP-Authentication-Scheme" required="false" max="1"/>
				<rule avp="SIP-Authenticate" required="false" max="1"/>
                <rule avp="SIP-Authorization" required="false" max="1"/>
                <rule avp="SIP-Authentication-Context" required="false" max="1"/>
                <rule avp="Confidentiality-Key" required="false" max="1"/>
                <rule avp="Integrity-Key" required="false" max="1"/>
                <rule avp="SIP-Digest-Authenticate" required="false" max="1"/>
                <rule avp="Framed-IP-Address" required="false" max="1"/>
                <rule avp="Framed-IPv6-Prefix" required="false" max="1"/>
                <rule avp="Framed-Interface-Id" required="false" max="1"/>
                <rule avp="Line-Identifier" required="false" />
			</data>
        </avp>
        <avp name="SIP-Item-Number" code="613" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="Unsigned32"/>
        </avp>
        <avp name="Server-Assignment-Type" code="614" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.15 -->
            <data type="Enumerated">
                <item code="0" name="NO_ASSIGNMENT"/>
                <item code="1" name="REGISTRATION"/>
                <item code="2" name="RE_REGISTRATION"/>
                <item code="3" name="UNREGISTERED_USER"/>
                <item code="4" name="TIMEOUT_DEREGISTRATION"/>
                <item code="5" name="USER_DEREGISTRATION"/>
                <item code="6" name="TIMEOUT_DEREGISTRATION_STORE_SERVER_NAME"/>
                <item code="7" name="USER_DEREGISTRATION_STORE_SERVER_NAME"/>
                <item code="8" name="ADMINISTRATIVE_DEREGISTRATION"/>
                <item code="9" name="AUTHENTICATION_FAILURE"/>
                <item code="10" name="AUTHENTICATION_TIMEOUT"/>
                <item code="11" name="DEREGISTRATION_TOO_MUCH_DATA"/>
                <item code="12" name="AAA_USER_DATA_REQUEST"/>
                <item code="13" name="PGW_UPDATE"/>
                <item code="14" name="RESTORATION"/>
            </data>
        </avp>
        <avp name="Deregistration-Reason" code="615" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.16 -->
            <data type="Grouped">
				<rule avp="Reason-Code" required="true" max="1"/>
				<rule avp="Reason-Info" required="false" max="1"/>
			</data>
        </avp>
        <avp name="Reason-Code" code="616" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.15 -->
            <data type="Enumerated">
                <item code="0" name="PERMANENT_TERMINATION"/>
                <item code="1" name="NEW_SERVER_ASSIGNED"/>
                <item code="2" name="SERVER_CHANGE"/>
                <item code="3" name="REMOVE_S-CSCF"/>
            </data>
        </avp>
        <avp name="Reason-Info" code="617" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="UTF8String"/>
        </avp>
        <avp name="Charging-Information" code="618" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.20 -->
            <data type="Grouped">
				<rule avp="Primary-Event-Charging-Function-Name" required="false" max="1"/>
				<rule avp="Secondary-Event-Charging-Function-Name" required="false" max="1"/>
				<rule avp="Primary-Charging-Collection-Function-Name" required="false" max="1"/>
				<rule avp="Secondary-Charging-Collection-Function-Name" required="false" max="1"/>
            </data>
        </avp>
        <avp name="Primary-Event-Charging-Function-Name" code="619" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="DiameterURI"/>
        </avp>
        <avp name="Secondary-Event-Charging-Function-Name" code="620" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="DiameterURI"/>
        </avp>
        <avp name="Primary-Charging-Collection-Function-Name" code="621" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="DiameterURI"/>
        </avp>
        <avp name="Secondary-ChargingCollection-Function-Name" code="622" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="DiameterURI"/>
        </avp>

        <avp name="User-Authorization-Type" code="623" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.24 -->
            <data type="Enumerated">
                <item code="0" name="REGISTRATION"/>
                <item code="1" name="DE_REGISTRATION"/>
                <item code="2" name="REGISTRATION_AND_CAPABILITIES"/>                
            </data>
        </avp>
        <avp name="User-Data-Already-Available" code="624" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="Enumerated">
                <item code="0" name="USER_DATA_NOT_AVAILABLE"/>
                <item code="1" name="USER_DATA_ALREADY_AVAILABLE"/>
            </data>
        </avp>
        <avp name="Confidentiality-Key" code="625" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="OctetString"/>
        </avp>
        <avp name="Integrity-Key" code="626" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="OctetString"/>
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


        <avp name="Associated-Identities" code="632" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <!-- 3GPP TS 29.229 Section 6.3.29 -->
            <data type="Grouped">
                <rule avp="User-Name" required="false"/>
            </data>
        </avp>

        <avp name="SIP-Digest-Authenticate" code="635" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <!-- 3GPP TS 29.229 Section 6.3.29 -->
            <data type="Grouped">
                <rule avp="Digest-Realm" required="true" max="1"/>
                <rule avp="Digest-Algorithm" required="false" max="1"/>
                <rule avp="Digest-QoP" required="true" max="1"/>
                <rule avp="Digest-HA1" required="true" max="1"/>
            </data>
        </avp>

        <avp name="UAR-Flags" code="637" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32" />
        </avp>
        <avp name="Loose-Route-Indication" code="638" must="V"  must-not="M"  may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="Enumerated">
                <item code="0" name="LOOSE_ROUTE_NOT_REQUIRED"/>
                <item code="1" name="LOOSE_ROUTE_REQUIRED"/>
            </data>
        </avp>
        <avp name="SCSCF-Restoration-Info" code="639" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <!-- 3GPP TS 29.229 Section 6.3.46 -->
            <data type="Grouped">
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Restoration-Info" required="true" min="1"/>
                <rule avp="SIP-Authentication-Scheme" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Path" code="640" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Contact" code="641" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Subscription-Info" code="642" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Call-ID-SIP-Header" required="true" max="1"/>
                <rule avp="From-SIP-Header" required="true" max="1"/>
                <rule avp="To-SIP-Header" required="true" max="1"/>
                <rule avp="Record-Route" required="true" max="1"/>
                <rule avp="Contact" required="true" max="1"/>
            </data>
        </avp>

        <avp name="Call-ID-SIP-Header" code="643" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="From-SIP-Header" code="644" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="To-SIP-Header" code="645" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Record-Route" code="646" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>
        <avp name="Associated-RegisteredIdentities" code="647" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="User-Name" required="false"/>
            </data>
        </avp>
        <avp name="Multiple-Registration-Indication" code="648" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="NOT_MULTIPLE_REGISTRATION"/>
                <item code="1" name="MULTIPLE_REGISTRATION"/>
            </data>
        </avp>
        <avp name="Restoration-Info" code="649" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="Path" required="true" max="1"/>
                <rule avp="Contact" required="true" max="1"/>
                <rule avp="Initial-CSeq-Sequence-Number" required="false" max="1"/>
                <rule avp="Call-ID-SIP-Header" required="false" max="1"/>
                <rule avp="Subscription-Info" required="false" max="1"/>
            </data>
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

        <avp name="LIR-Flags" code="653" must="V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="Initial-CSeq-Sequence-Number" code="654" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
        <avp name="SAR-Flags" code="655" must="V" may="P" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>
    </application>
</diameter>`
