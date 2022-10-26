# lego-ui

lego-ui is an graphical user interface written in go developed on top of the [go-acme/lego](https://github.com/go-acme/lego) library to make it simple for everyone to get 
LetsEncrypt certificates using the dns-01 acme challange

# Vision

For novices it can be pretty difficult to generate new letsencrypt certificates. While there are some great solutions like from [Synology](https://kb.synology.com/de-de/DSM/tutorial/How_to_enable_HTTPS_and_create_a_certificate_signing_request_on_your_Synology_NAS) or [Cpanel](https://blog.cpanel.com/how-to-configure-and-manage-lets-encrypt-in-cpanel/), or even GitHub, which are great to use but force the user to host their content in certain (propitary) environments; we loose the ability to host software on their own devices. If we want decentralised self-hosted applications to become the norm, we have to invest into makeing them easy to set up.

Now currently if users really want to do it all by themselves, the usually need to traverse into the command line
