# ========================================================================================================
#
# PROJECT REALITY SERVER COMMON SETTINGS DEFINITION
#
# This file can be edited by any server (public or private).
# 
# $Id: realityconfig_common.py 34646 2021-02-24 15:16:03Z prbot $
#
#
# PROJECT REALITY LOCALIZATION
# Edit this setting to set the language file to be used for the in-game messages
# These files are located in mods/pr/localization/language/pr.utxt
# All strings starting with PYTHON_ are used by the python code messages
# Default is english
PRL10N = 'english'
#
#
# PROJECT REALITY MUTINY
# Edit this setting to set the use of PR's own mutiny system
# based on only Squad Leaders votes
# Default is enabled
PRMUTINY = 1
#
#
# PROJECT REALITY STATS CONSTANTS
# Edit this setting to set the use of PR's own constants file instead of the one 
# defined in /stats/constants.py - if you use custom stats, turn this off
# Default is enabled
PRCONSTANTS = 1
#
#
# PROJECT REALITY TIME LIMIT
# Edit this setting to set the time limit for all game modes expressed in seconds
# The purpose of a time limit is just to provide an in game clock
# Default is 14400 (4 hours). Above 4 hours BF2CC will throw an error message
# Set to 0 to disable time limit
PRTIMELIMIT = 14400
#
#
# PROJECT REALITY ROUND START DELAY
# Edit this setting to set the starting delay at the beginning of a game
# The purpose of a start delay is to allow squads to be made, and assets to be claimed
# Default is 240 (4 minutes), minimum of 120 (2 minutes), maximum of 300 (5 minutes) (For public servers)
PRROUNDSTARTDELAY = 240
#
#
# PROJECT REALITY BOT - PRBOT SPECTATOR CAMERA
# Edit this setting to set the player names (without prefixes and all lowercase) 
# that can spawn the spectator camera in-game with the console command "rcon prbot"
# while the server is running private, local or coop configs.
# Example: PRSPECTATORS = [ 'username', 'otherusername' ]
PRSPECTATORS = []


# Override server-side timeout
# How long (in milliseconds) does a server keep clients when no packets have been received from them
# None for default (45000ms)
PRCONNECTIONTIMEOUT = None



# Maximum commands from client to buffer.
# BF2's value is 4, and is not needed with modern connections.
# possible values are 2 - 4
# Lower value create mucg less latency for clients, but may cause inputs to be ignored when ping is unstable (bad WiFi)
ACTIONBUFFER_MAXSIZE = 2



# Maximum outgoing packet size
# Max value is 0x540 and is recommended.
# Players with bandwidth problems may independently set "GeneralSettings.setConnectionType 4" to request smaller packets, but I doubt there's anyone who can't handle this max size
OUTGOING_PACKET_MAXSIZE = 0x540



#
#
# END OF EDITABLE PARAMETERS
