
#!/usr/bin/bash

if [ -z "$1" ]
	echo "Usage: ./recon.sh <ip>"
	exit 1
fi

printf "\n---- NMAP ----\n\n" > results

echo "Running Nmap..."
nmap $1 | tail -n +5 \ head -n -3 >>results
