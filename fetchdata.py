"""
    parse_wikitable.py

    MediaWiki Action API Code Samples
    Demo of `Parse` module: Parse a section of a page, fetch its table data and save
    it to a CSV file

    MIT license
"""

import csv
import requests

S = requests.Session()

URL = "https://en.wikipedia.org/w/api.php"

TITLE = "Game_of_Thrones_(season_8)"

PARAMS = {
    'action': "parse",
    'page': TITLE,
    'prop': 'wikitext',
    'section': 1,
    'format': "json"
}


def get_table():
    """ Parse a section of a page, fetch its table data and save it to a CSV file
    """
    res = S.get(url=URL, params=PARAMS)
    data = res.json()
    wikitext = data['parse']['wikitext']['*']
    lines = wikitext.split('|')
    entries = []

    for line in lines:
        line = line.strip()
        entries.append(line)

    file = open("Game_of_Thrones_.csv", "w")
    writer = csv.writer(file)
    writer.writerows(entries)
    file.close()

if __name__ == '__main__':
    get_table()