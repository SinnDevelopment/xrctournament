# xRC Tournament

Tournament Software for FRC or FTC Simulator Competitions.

Built as a "one-stop shop" for a tournament, including a leaderboard, qualification and elimination match support.

This is intended to compliment the existing AWS tutorial, by running a tournament completely autonomously, with the option of being headless on a remote server.

## Features

* When a match is completed that has the same lineup as a scheduled match, that match is marked as completed. This allows matches to be completed out of order based on people's availability. The scores and individual contributions are automatically saved, without needing any user interaction.
* The ability for multiple stages of matches to be run (practice, qualification, playoff) by simply renaming the `matches.json` file to something else, and loading a different `schedule.csv` file.
* Live leaderboard, including user standings, and match results.

## Planned Features

* Alliance Selection Schedule Tool - Generate a schedule based on an alliance selection input - Could be static JS.
