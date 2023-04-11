# dead-or-alive-golang
A service that I made to test printers across my works network, takes CSV of unique third octets to add to the "madeIP" and reports via a log file. Uses cmd ping output to generate results.

The octets are passed in with CSV in this format:
Office Name [40],
Office Name [3],
Office Name [71],
Office name [99],
