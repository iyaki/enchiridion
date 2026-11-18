---
title: "OAuth 2.0 for Dummies"
notion_id: 780e8bdd-db5e-448a-8623-e511c9760a56
notion_url: https://app.notion.com/p/OAuth-2-0-for-Dummies-780e8bdddb5e448a8623e511c9760a56
last_edited: 2023-01-13T16:19:00.000Z
source_url: https://hackernoon.com/oauth-20-for-dummies
tags: ["English", "Information Security", "System Design / Software Architecture", "Article", "Hackernoon"]
---
## Too Long; Didn't Read

OAuth is an open-standard authorization protocol that lets a service use another service without requiring the security details (username, password, etc.) of the user. OAuth lets you authorize one application to access your data, or use features in another application on your behalf, without giving them your password. It is a security standard where you give one application permission to access *your data* in another application. Let’s dive into OAuth to learn more about the differences between authentication and authorization.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Maneesha Indrachapa**

An enthusiastic engineering graduate from the Computer Sc...

_OAuth 2.0 is the industry-standard protocol for authorization._

You can get the above definition if you Googled “what is OAuth”.

_But why OAuth in the first place?_ In the early days of the Internet, sharing information between services was easy. You simply gave your username and password for one service to another so they could sign in to your account and grab whatever information they wanted.

With the development of the internet and the services being digitalized, there is a problem sharing your username and password to service because that service can be developed using malicious intent.

So to get rid of those kinds of scenarios OAuth was defined. As the OAuth definition states, OAuth is an open-standard authorization protocol that lets a service use another service without requiring the security details (username, password, etc.) of the user.

When talking about Authorization it is always people who tend to mix up with authentication. Let’s look are the differences between authentication and authorization.

## Authentication vs Authorization

Let’s see the differences between Authentication and Authorization,

- 

Authentication verifies who the user is. Authorization determines what resources a user can access.

- 

Authentication works through passwords, one-time pins, biometric information, and other information provided or entered by the user. Authorization works through settings that are implemented and maintained by the organization.

- 

Authentication is the first step of a good identity and access management process. Authorization always takes place after authentication.

- 

Authentication is visible to and partially changeable by the user. Authorization isn’t visible to or changeable by the user.

So now it is clear about Authentication and Authorization is and have a basic idea of why we need a standard protocol to give one application permission to access _your data_ in another application let's dive into OAuth.

## Let’s Dive into OAuth

Basically, [**OAuth 2.0**](https://oauth.net/2/?ref=hackernoon.com)** is a security standard where you give one application permission to access **_**your data**_** in another application. The steps to grant permission, or **_**consent**_**, are often referred to as **_**authorization**_** or even **_**delegated authorization**_**.** You authorize one application to access your data, or use features in another application on your behalf, without giving them your password.

Consider the most popular example, the valet key for your car. The vehicle owner’s key can control everything in the car such as starting it, opening doors and windows, accessing the glove box, opening the trunk of the car etc. But the valley key can only be used to start the car and lock/unlock the doors. _This is the concept behind OAuth. Providing a key with limited access rights_.

Before going into detail about how OAuth Works lets get familiar with some words used in,

- **Resource Owner:** You! You are the owner of your identity, your data, and any actions that can be performed with your accounts.
- **Client:** The application that wants to access data or perform actions on behalf of the Resource Owner.
- **Authorization Server:** The application that knows the Resource Owner, where the Resource Owner already has an account.
- **Resource Server:** The Application Programming Interface (API) or service the Client wants to use on behalf of the Resource Owner.
- **Access Token:** The key the client will use to communicate with the Resource Server. This is like a valet key or key card that gives the Client permission to request data or perform actions with the Resource Server on your behalf. An access token is a string representing an authorization issued to the client. The string is usually opaque to the client. Tokens represent specific scopes and durations of access, granted by the resource owner and enforced by the resource server and authorization server.
- **Refresh Token:** Refresh tokens are the credentials used to obtain access tokens. Refresh tokens are issued to the client by the authorization server and are used to obtain a new access token when the current access token becomes invalid or expires or to obtain additional access tokens with an identical or narrower scope.
- **Scope:** The level of access that the application is requesting. The authorization server may fully or partially ignore the scope requested by the client, based on the authorization server policies or the resource owner’s instructions.

Here, the resource owner (you) wants to create a Spotify account and use the profile picture and the profile details you have in the Facebook account used in the Spotify profile. The above diagram illustrates the flow of access delegation via OAuth. Here are the steps associated with the process

1 - The client asks for the resource owner’s permission to use the resource owner’s Facebook Profile Details(resource) which is hosted in a resource server and the resource owner accepts the request.

2 - Then, the client kindly requests access with a scope(To access Facebook profile details) to the service by contacting the Facebook authorization server. This process is called **Authorization Request**.

3- Next is obtaining the **Authentication and Consent of** the owner. In this step, the authorization server verifies the client and then contacts the resource owner asking to authenticate him or herself. Therefore, the owner has to log in to Facebook. Once the owner is authenticated, the authorization server will ask for the owner’s permission to allow the Profile picture, name, birthday and a few other details. This is called the **resource owner’s consent**. If the owner allows it the server will send a positive response to the client.

4 - When the authorization server gets owner consent, the server will send the client a key (some sort of passcode) to be used to access the resources at the resource server. This is known as an **Access Token**. Along with the access, the token authorization server sends a **Refresh Token** because access tokens have a small validity period when compared to Refresh Tokens, Once an access token has expired, the resource server will reject the token. Then the client will send the refresh token back to the authorization server and the authorization server will issue a new access token to the client.

5 - Now the client happily contacts the resource server, presenting the access token in both hands. This is called a **Resource Request.**

6 - Finally, for added protection, resource servers don’t just hand over the resources even though the client has an access token, it validates the access token with the authorization server. This is called **Token Introspection**. After validation, the client will get to access the resources. The client can now use the access token to access the resources anytime it wants without contacting the authorization server until the access token is expired.

7 - The client can now use the access token to access the resources anytime it wants without contacting the authorization server until the access token is expired.

As you can see in this process the Resource Owner has never provided his credentials to Spotify. The owner only shared the credentials with the Facebook authorization server and send an Access token + Refresh token to Spotify. So the user credentials of the user are safe with the user and never being exposed to outside parties. In a nutshell, this is what happens in OAuth in access delegation.

There are several ways to get Access token in OAuth and different ways to use it. They are called _**Grant types**_. They will be explained in another blog.

## References

- [https://developer.okta.com/blog/2019/10/21/illustrated-guide-to-oauth-and-oidc](https://developer.okta.com/blog/2019/10/21/illustrated-guide-to-oauth-and-oidc?ref=hackernoon.com)
- [https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth?ref=hackernoon.com)
- [https://wso2.com/blogs/thesource/2019/08/oauth-2-basics/](https://wso2.com/blogs/thesource/2019/08/oauth-2-basics/?ref=hackernoon.com)
- https://www.sailpoint.com/identity-library/difference-between-authentication-and-authorization

This article was originally published in https://maneeshaindrachapa.medium.com/oauth-simplified-921b41dbb6a8

L O A D I N G
. . . comments &
