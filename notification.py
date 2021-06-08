#!python3

import os

from pathlib import Path

third_party = Path('third-party')
notification = open('NOTIFICATION', 'wt')

for root, dirs, files in os.walk(third_party):
    if not files:
        continue

    root = Path(root)
    pkg = root.relative_to(third_party)
    notification.write(f'{pkg} (http://{pkg})\n\n')

    for f in files:
        file = root / f
        contents = file.read_text()
        notification.write(f + '\n\n' + contents + '\n')

notification.write('\n')
notification.close()
