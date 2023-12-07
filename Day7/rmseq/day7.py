import sys
from collections import Counter

values_pt1 = {'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10}
values_pt2 = {'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10}


def parse(file):
    for line in file:
        hand, bet = line.split()
        yield Counter(hand), hand, int(bet)


def hand_ranking(counter):
    return [count[1] for count in counter.most_common()[:2]]


def card_values(hand, values):
    return [int(card) if card not in values else values[card] for card in hand]


def solve(data):
    def calc(part):
        res = 0
        for i, h in enumerate(sorted(part, key=lambda x: (x[1], x[0]))):
            res += (i + 1) * h[2]
        return res

    pt1, pt2 = list(), list()
    for counter, hand, bid in data:
        pt1.append((card_values(hand, values_pt1), hand_ranking(counter), bid))

        if 'J' in counter and len(counter) > 1:
            mc = counter.most_common()
            target = mc[0][0]
            if target == 'J':
                target = mc[1][0]
            counter[target] += counter.pop('J')

        pt2.append((card_values(hand, values_pt2), hand_ranking(counter), bid))

    return calc(pt1), calc(pt2)


if len(sys.argv) != 2:
    print(f"usage: {sys.argv[0].split("/")[-1]} <path>")
    exit(1)

with open(sys.argv[1], 'r') as f:
    print(*solve(parse(f)))
