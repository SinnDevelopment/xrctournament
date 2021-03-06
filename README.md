# xRC Tournament

Tournament Software for FRC or FTC Simulator Competitions.

Built as a "one-stop shop" for a tournament, including a leaderboard, qualification and elimination match support.

This is intended to compliment the existing AWS tutorial, by running a tournament completely autonomously, with the
 option of being headless on a remote server.

## Features

* When a match is completed that has the same lineup as a scheduled match, that match is marked as completd. This allows
 matches to be completed out of order based on people's availability. The scores and individual contributions are 
 automatically saved, without needing any user interaction.
* The ability for multiple stages of matches to be run (practice, qualification, playoff) by simply renaming the
 `matches.json` file to something else, and loading a different `schedule.csv` file.
* Live leaderboard, including user standings, and match results.

## Planned Features

* Alliance Selection Schedule Tool - Generate a schedule based on an alliance selection input - Could be static JS.

## Usage

Download the appropriate binary for your OS from the releases page. You can either run the binary once to generate the 
`config.json` file, or copy the existing template in the repository.
After the config file is generated, edit to your tournament's liking, then relaunch `xrctournament`.

### Schedule Importing
xrctournament can handle quals, and elim schedule. Set the filepath for the given schedule you want to import.
Qualification matches cannot be imported at the same time as elimination matches due to the fact that both overwrite the
 master match list. This may be handled by an "active" level like FMS, but that's not planned.

### Match Results
In non-scheduled mode, match results are automatically registered to the master list, and exported to `matches.json` after each match.


### FRC.Events sync style to cloud