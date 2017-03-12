#!/bin/bash
echo <<EOF
{ 
	"notebooks": [
		{"id": 0, "title": "queries", "deleted": false},
		{"id": 1, "title": "oneliners", "deleted": false},
		{"id": 2, "title": "configs", "deleted": false},
		{"id": 3, "title": "workflows", "deleted": false},
		{"id": 4, "title": "randoms", "deleted": true}
	],
	"notes": [
EOF

for i in {1..1024};do
   echo {\"id\":$i, \"notebook\": $((i%5)), \"title\": \"note \#$i\", \"body\": \"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum\"},
done

echo <<EOF
	]
}
EOF
