#!/usr/bin/env python3
"""
Imports a old genesis' cards into cardchain
get_card_content has to be build first
"""

import json
import sys
import ctypes

lib = ctypes.cdll.LoadLibrary('./get_card_content.so')

def main(card_records):
    artist = "Cooler Artist".encode("utf-8")
    creator = "Cooler Typ".encode("utf-8")
    for i, card in enumerate([c for c in card_records if c["Content"] not in ["", "e30="] and c["Image"] != ""]):
        lib.make_buy_card_scheme_request(creator, "800000000000ucredits".encode("utf-8"))

        content = card["Content"].encode('utf-8')
        notes = card['Notes'].encode("utf-8")
        id = i+1
        lib.make_save_card_content_request(creator, i+1, content, notes, artist)

        artwork = card["Image"].encode("utf-8")
        full_art = card["FullArt"]
        lib.make_add_artwork_request(artist, i+1, artwork, full_art)


if __name__ == "__main__":
    with open(sys.argv[1], "r") as file:
        card_records = json.load(file)["app_state"]["cardservice"]["card_records"]

    main(card_records)
