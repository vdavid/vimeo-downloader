Source: https://bojtiandrea.hu/kistestver-szuletik-zart-kurzusfelulet/

## Usage
1. Go to page
2. Go to Dev Tools → Network → Fetch/xhr, filter for "master"
3. Hover through all videos in order to load them
4. Copy the URL of each request (quite tedious, I know), into the input/urls.txt. Mind that they'll expire in an hour.
5. Select, then copy the complete page content and paste it to chatgpt, ask it to convert it to a CSV with 4 columns: 1. Index, 2. Filename (inferred from the title, but somewhat shorter), prefixed with the index, all lowercase, separated by underscore, and with the extension "mp4", 3. Title, without "fejezet", and no all caps, 4. Description (can be multiline)
6. Get the saved CSV and save it as input/chapters.csv, without any header.
7. Copy-paste your cookies into the input/cookies.txt
8. Copy-paste your headers for any of the "master" requests into the input/headers.txt
9. Run the script
10. Optionally, go to the Network tab again, and search for "vtt", then download the subtitles one by one. I haven't found a way to automate this, or even to find their correct order automatically, so it's a bit of a mess.