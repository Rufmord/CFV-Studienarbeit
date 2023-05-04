echo "Laufende Prozesse"
ps | grep "sleep"
echo "Offene eingehende Ports"
lsof -i -P | grep -i "listen"
echo "Offene ausgehende Ports"
lsof -i -P | grep "thi.de"
echo "Offene Dateien"
lsof | grep open_me.txt
