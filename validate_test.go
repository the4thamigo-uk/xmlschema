package xmlschema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateMetaData(t *testing.T) {
	data := []byte(`
<md:EntityDescriptor entityID="https://sp.example.com/shibboleth"
                     xmlns:ds="http://www.w3.org/2000/09/xmldsig#"
                     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                     xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"
                     xmlns:mdui="urn:oasis:names:tc:SAML:metadata:ui">

   <md:SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
      <md:Extensions>
         <init:RequestInitiator
            xmlns:init="urn:oasis:names:tc:SAML:profiles:SSO:request-init"
            Binding="urn:oasis:names:tc:SAML:profiles:SSO:request-init"
            Location="https://sp.example.com/Shibboleth.sso/Login" />
         <idpdisc:DiscoveryResponse
            xmlns:idpdisc="urn:oasis:names:tc:SAML:profiles:SSO:idp-discovery-protocol"
            Binding="urn:oasis:names:tc:SAML:profiles:SSO:idp-discovery-protocol"
            Location="https://sp.example.com/Shibboleth.sso/DS" index="1" />

         <mdui:UIInfo>
            <mdui:DisplayName xml:lang="en">GARR Test SP</mdui:DisplayName>
            <mdui:DisplayName xml:lang="it">GARR SP di Test</mdui:DisplayName>
            <mdui:Description xml:lang="en">This is a Service Provider useful for testing</mdui:Description>
            <mdui:Description xml:lang="it">Questo è un Service Provider utile per i test</mdui:Description>
            <mdui:InformationURL xml:lang="en">https://sp.example.com/en/information.html</mdui:InformationURL>
            <mdui:InformationURL xml:lang="it">https://sp.example.com/it/information.html</mdui:InformationURL>

            <mdui:Logo height="16" width="16" xml:lang="en">https://sp.example.com/en/images/communityLogo-16x16.png</mdui:Logo>
            <mdui:Logo height="16" width="16" xml:lang="it">https://sp.example.com/it/images/communityLogo-16x16.png</mdui:Logo>
            <mdui:Logo height="60" width="80" xml:lang="en">https://sp.example.com/en/images/communityLogo-80x60.png</mdui:Logo>
            <mdui:Logo height="60" width="80" xml:lang="it">https://sp.example.com/it/images/communityLogo-80x60.png</mdui:Logo>

            <mdui:PrivacyStatementURL xml:lang="en">https://sp.example.com/en/privacyStatement.html</mdui:PrivacyStatementURL>
            <mdui:PrivacyStatementURL xml:lang="it">https://sp.example.com/it/privacyStatement.html</mdui:PrivacyStatementURL>
         </mdui:UIInfo>
      </md:Extensions>

      <md:KeyDescriptor>
         <ds:KeyInfo>
            <ds:X509Data>
               <ds:X509Certificate>
                  MIICajCCAdOgAwIBAgIBADANBgkqhkiG9w0BAQQFADBSMQswCQYDVQQGEwJ1czETMBEGA1UECAwKQ2FsaWZvcm5pYTEVMBMGA1UECgwMT25lbG9naW4gSW5jMRcwFQYDVQQDDA5zcC5leGFtcGxlLmNvbTAeFw0xNDA3MTcwMDI5MjdaFw0xNTA3MTcwMDI5MjdaMFIxCzAJBgNVBAYTAnVzMRMwEQYDVQQIDApDYWxpZm9ybmlhMRUwEwYDVQQKDAxPbmVsb2dpbiBJbmMxFzAVBgNVBAMMDnNwLmV4YW1wbGUuY29tMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC7vU/6R/OBA6BKsZH4L2bIQ2cqBO7/aMfPjUPJPSn59d/f0aRqSC58YYrPuQODydUABiCknOn9yV0fEYm4bNvfjroTEd8bDlqo5oAXAUAI8XHPppJNz7pxbhZW0u35q45PJzGM9nCv9bglDQYJLby1ZUdHsSiDIpMbGgf/ZrxqawIDAQABo1AwTjAdBgNVHQ4EFgQU3s2NEpYx7wH6bq7xJFKa46jBDf4wHwYDVR0jBBgwFoAU3s2NEpYx7wH6bq7xJFKa46jBDf4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQQFAAOBgQCPsNO2FG+zmk5miXEswAs30E14rBJpe/64FBpM1rPzOleexvMgZlr0/smF3P5TWb7H8Fy5kEiByxMjaQmml/nQx6qgVVzdhaTANpIE1ywEzVJlhdvw4hmRuEKYqTaFMLez0sRL79LUeDxPWw7Mj9FkpRYT+kAGiFomHop1nErV6Q==
							</ds:X509Certificate>
            </ds:X509Data>
         </ds:KeyInfo>
      </md:KeyDescriptor>

      <md:ArtifactResolutionService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
         Location="https://sp.example.com/Shibboleth.sso/Artifact/SOAP"
         index="1" />

      <md:SingleLogoutService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
         Location="https://sp.example.com/Shibboleth.sso/SLO/SOAP" />
      <md:SingleLogoutService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
         Location="https://sp.example.com/Shibboleth.sso/SLO/Redirect" />
      <md:SingleLogoutService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
         Location="https://sp.example.com/Shibboleth.sso/SLO/POST" />
      <md:SingleLogoutService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Artifact"
         Location="https://sp.example.com/Shibboleth.sso/SLO/Artifact" />

      <md:ManageNameIDService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
         Location="https://sp.example.com/Shibboleth.sso/NIM/SOAP" />
      <md:ManageNameIDService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
         Location="https://sp.example.com/Shibboleth.sso/NIM/Redirect" />
      <md:ManageNameIDService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
         Location="https://sp.example.com/Shibboleth.sso/NIM/POST" />
      <md:ManageNameIDService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Artifact"
         Location="https://sp.example.com/Shibboleth.sso/NIM/Artifact" />

      <md:AssertionConsumerService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
         Location="https://sp.example.com/Shibboleth.sso/SAML2/POST"
         index="1" />
      <md:AssertionConsumerService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST-SimpleSign"
         Location="https://sp.example.com/Shibboleth.sso/SAML2/POST-SimpleSign"
         index="2" />
      <md:AssertionConsumerService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Artifact"
         Location="https://sp.example.com/Shibboleth.sso/SAML2/Artifact"
         index="3" />
      <md:AssertionConsumerService
         Binding="urn:oasis:names:tc:SAML:2.0:bindings:PAOS"
         Location="https://sp.example.com/Shibboleth.sso/SAML2/ECP"
         index="4" />

      <md:AttributeConsumingService index="1">
        <md:ServiceName xml:lang="en-US">string</md:ServiceName> 

         <!-- example for the required attribute: mail -->
         <md:RequestedAttribute FriendlyName="mail"
            Name="urn:oid:0.9.2342.19200300.100.1.3"
            NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
            isRequired="true" />

         <!-- example for the required attribute: eduPersonPrincipalName -->
         <md:RequestedAttribute FriendlyName="eppn"
            Name="urn:oid:1.3.6.1.4.1.5923.1.1.1.6"
            NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
            isRequired="true" />

      </md:AttributeConsumingService>
   </md:SPSSODescriptor>

   <md:Organization>
      <md:OrganizationName xml:lang="en">Consortium GARR</md:OrganizationName>
      <md:OrganizationName xml:lang="it">Consortium GARR</md:OrganizationName>

      <md:OrganizationDisplayName xml:lang="en">Consortium GARR</md:OrganizationDisplayName>
      <md:OrganizationDisplayName xml:lang="it">Consortium GARR</md:OrganizationDisplayName>

      <md:OrganizationURL xml:lang="en">http://www.garr.it/b/eng</md:OrganizationURL>
      <md:OrganizationURL xml:lang="it">https://www.garr.it</md:OrganizationURL>
   </md:Organization>

   <md:ContactPerson contactType="technical">
      <md:EmailAddress>mailto:example.technical.contact@garr.it</md:EmailAddress>
   </md:ContactPerson>
   <md:ContactPerson contactType="support">
      <md:EmailAddress>mailto:example.support.contact@garr.it</md:EmailAddress>
   </md:ContactPerson>

</md:EntityDescriptor>`)
	err := Validate(data, Metadata)
	assert.Nil(t, err)
}

func TestValidateProtocol(t *testing.T) {
	data := []byte(`
<samlp:AuthnRequest xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol" xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion" ID="ONELOGIN_809707f0030a5d00620c9d9df97f627afe9dcc24" Version="2.0" ProviderName="SP test" IssueInstant="2014-07-16T23:52:45Z" Destination="http://idp.example.com/SSOService.php" ProtocolBinding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" AssertionConsumerServiceURL="http://sp.example.com/demo1/index.php?acs">
  <saml:Issuer>http://sp.example.com/demo1/metadata.php</saml:Issuer>
  <samlp:NameIDPolicy Format="urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress" AllowCreate="true"/>
  <samlp:RequestedAuthnContext Comparison="exact">
    <saml:AuthnContextClassRef>urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport</saml:AuthnContextClassRef>
  </samlp:RequestedAuthnContext>
</samlp:AuthnRequest>`)
	err := Validate(data, Protocol)
	assert.Nil(t, err)
}

func TestValidateXHTML(t *testing.T) {
	data := []byte(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <title>Title of document</title>
</head>
<body>
  <p>some content</p>
</body>
</html>`)
	err := Validate(data, XHTML)
	assert.Nil(t, err)
}

func TestValidateParseError(t *testing.T) {
	data := []byte(`invalid`)
	err := Validate(data, Metadata)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to create parse input: failed to document: Entity: line 1: parser error : Start tag expected, '<' not found", err.Error())
}
