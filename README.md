# dynamic-port-scanning-in-GO-lang
This is simple repositry use to detect which port is open. It is a custom tool built in GO LANG. <br>
This is CUSTOM tool for those who want to understand how port scanning works in go lang and how can custom port scanning tool can be created.
<h3>
  Using
  </h3>
  <p> It dials the port with address using <b> net.Dial("tcp", "stackoverflow.com")</b>.
</p>
<h3>
  How to use it
  </h3>
  <ol> 
  <li>To scan ports from 0-1024 use .\pscan.exe stackoverflow.com</li>
  <li>If you want to increase the number of ports to scan change 	var maxportnumber = 1024 from <i>pscan.go</i> and rebuild the project.</li>
  <li>To scan specific port use .\single-port-scan.exe stackoverflow.com:443</li>

</ol>
  
  <b>Special thanks to booK:Black-Hat-Go_Go-Programming-For-Hackers-and-Pentesters</b>
