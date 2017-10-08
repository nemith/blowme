# blowme
A quicky and dirty program to control my sprinklers for winterizing them.

I have a an old IrrigationCaddy S1 which as a sort of REST interface for controlling it (discovered with the webUI and issuing commands and seeing what was on the wire).  So I can easily start and stop zones.

This just loops through all the zones I have, turns on the zone for 30 seconds and then allows enough time for my compressor to recharge the tank before moving onto the next zone.