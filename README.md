# discourse-user-api-key

### Purpose
With a User-API-Key it is possible to perform certain actions automatically in a (Discourse) forum. The User-API-Key is bound to a user. The actions are performed on behalf of the user for whom the User-API-Key was created.

This utility allows to request a User-API-Key from a (Discourse) forum. The utility is available as command line interface application for different operating systems.

### Binaries
Precompiled binaries for many operating systems can be found here: **Releases -> Assets**

### Usage
``` text
% ./discourse-user-api-key -h

Program:
  Name    : discourse-user-api-key
  Release : v1.0.0 - 2022/10/31
  Purpose : obtain Discourse User-API-Key
  Info    : This program obtains a User-API-Key for a Discourse forum.

Usage:
  ./discourse-user-api-key -forum=string [-application=string] [-client=string] [-scopes=list] [-nonce=string] [-verbose]

Examples:
  ./discourse-user-api-key -forum=community.openstreetmap.org
  ./discourse-user-api-key -forum=meta.discourse.org -application=UltimateReaderWriter -scopes=read,write

Options:
  -application string
    	name of application shown on forum site (default "GenericDiscourseReader")
  -client string
    	client ID (default [generated unique UUID4])
  -forum string
    	Discourse forum URL
  -nonce string
    	random string generated once (default [generated URL-safe random string])
  -scopes string
    	comma-separated list of access scopes allowed for the key (default "read")
  -verbose
    	verbose output (maybe helpful in case of problems)

Workflow for getting an User-API-Key:
  Step 1: copy forum URL into your browser
  Step 2: authorize application access on forum site
  Step 3: copy encrypted User-API-Key data from forum site in here
  Step 4: save User-API-Key into your key vault
```

### Example
``` text
% ./discourse-user-api-key -forum=community.openstreetmap.org

Program:
  Name    : discourse-user-api-key
  Release : v1.0.0 - 2022/10/31
  Purpose : obtain Discourse User-API-Key
  Info    : This program obtains a User-API-Key for a Discourse forum.

Workflow for getting an User-API-Key:
  Step 1: copy forum URL into your browser
  Step 2: authorize application access on forum site
  Step 3: copy encrypted User-API-Key data from forum site in here
  Step 4: save User-API-Key into your key vault

Step 1: copy forum URL into your browser ...

https://community.openstreetmap.org/user-api-key/new?application_name=GenericDiscourseReader&client_id=47d3c48e-b6b4-4f5b-8822-6203aaa9ed97&nonce=3J9rujWeEQP_IEwasM-r6HN-sqY%3D&public_key=-----BEGIN+PUBLIC+KEY-----%0AMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo26hAwBt8dnQ93nwpFBw%0AX0KoxuskYvt7n6RN5l6NUKWsCTdC%2BGf%2Ftu2VvgTBPfgxDIfun2uIvF6n743m10qB%0AieCm8VIdPK2Yc3h4h58KfkZ4G9s%2BbJeLLv3g7uDCNMG8A2%2F8u5XSCxN%2FeuIkstvr%0AAmTBClrInJrq6pUELpgyoyZbNl1DpUl%2BX92%2FfUeGgdORHiKorAbZT4PCC5HE72pG%0AAdv45i6EvCHUfuISkB56DK4CduGH0v5dzZTXlnFi6iaNpH0pd8rJTvRCmrV9VutG%0A96fv%2B1MutjwfWaL90WJ2JwF%2FEVWfBe4cQATCVqRN5lQ79tdY5UwgPuyXS%2Big9xuB%0ALwIDAQAB%0A-----END+PUBLIC+KEY-----%0A&scopes=read

Step 2: authorize application access on forum site ...

Step 3: copy encrypted User-API-Key data from forum site in here (and press Enter) ...

Pp7IxXlZ4Yb6nCnQfJ5htMaffA9TAaH9RztOUYvI2IssBxZXWOQ5+LOJRtqO sTWJYAUFrIwwYPhJGq5Dc4oA7VFxy8cOVDfq3VkDNw5tA1T3OGeysiIjyRC8 KbzDZdS8xYhFdIyI9rYzhbnYeCazY/33LG/4UFf+Ld1nkx0GtwLGi8D75+sK 6ECnw/oDwGMoUFLjlQ4p47mOa0bDGh7xUS2h3VO7xee13ctvj873dqRmiPuN xaTsQXtf3v9EwVAY1Hj+xqrex78yj8U3hmyEXqssoGhNn8QlVAmn8E93dcdu TeSf9fhhgSoePc1a/kiEdRTlGVMcR5+WDqfNg1wDPw==

Decrypted User-API-Key data = {"key":"aa141fc6a7362231178d8d8e5609e768","nonce":"3J9rujWeEQP_IEwasM-r6HN-sqY=","push":false,"api":4}

User-API-Key = aa141fc6a7362231178d8d8e5609e768
-----------------------------------------------

Step 4: save User-API-Key into your key vault
```
