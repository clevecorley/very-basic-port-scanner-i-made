#Basic port scanner

A simple port scanner written in Go to check if ports on a target are open.
How It Works
  The program asks for the number of ports to scan and the target IP/domain.
  It attempts to connect to each port in the range and prints if it's open or closed.
  
#example out-put
  welcome to basic port scanner
  enter how much ports you want to scan:
  1000
  Target:
  127.0.0.1
  1 status: Open
