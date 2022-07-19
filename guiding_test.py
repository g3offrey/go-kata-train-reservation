"""
This is a guiding test for this kata. You can run it using python 3.3 like this:

    python3 -m unittest guiding_test.py

You need only need implement one of:
- the http interface
- the command line interface

Edit this file to remove one of the test cases as appropriate. 

If you are using the command line interface, you should also change the value of the 
"interpreter" and "reservation_script" variables to match the name of your command line program.
"""

import json
import subprocess
import unittest
import os

reservation_script = os.path.join("./reserve.sh")

class TrainReservationTest(unittest.TestCase):
    def test_reserve_seats_via_cmd(self):
        response = subprocess.check_output([reservation_script, "express_2000", "4"], stderr=subprocess.STDOUT, universal_newlines = True)
        reservation = json.loads(response)
        
        assert "express_2000" == reservation["train_id"]
        assert 4 == len(reservation["seats"])
        assert "1A" == reservation["seats"][0]
        assert "75bcd15" == reservation["booking_reference"]
