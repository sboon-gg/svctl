# ========================================================================================================
#
# PROJECT REALITY SERVER SETTINGS DEFINITION FOR PRIVATE SERVERS
#
# This file can be fully edited and it's automatically used by passworded internet servers
#
# This config is identical to the public one.
#
# $Id: realityconfig_private.py 38345 2023-06-28 07:58:37Z prbot $
#
#
import realityconstants as CONSTANTS

#
#
C = { }
#
#
# PROJECT REALITY TIME LIMIT
# Edit this setting to set the time limit for all game modes expressed in seconds.
# Default is 14400 (4 hours) defined in realityconfig_common.py PRTIMELIMIT variable.
# If you change this setting to anything other than None it will be used instead.
C['PRTIMELIMIT'] = None
#
#
# PROJECT REALITY ROUND START DELAY
# Edit this setting to set the start delay for all game modes
# Default is 180 seconds in realityconfig_common.py STARTDELAY variable.
# If you change this setting to anything other than 0 it will be used instead.
C['STARTDELAY'] = 0
#
#
# PROJECT REALITY WOUNDED TIME
# Edit these settings to determine the amount of seconds the player will stay wounded
# Default is 300 seconds
C['WOUNDED_TIME'] = 300
#
#
# PROJECT REALITY DEAD TIME
# Edit these settings to determine the amount of seconds the player will stay dead
# Default is 30 seconds
#
# Note that this time is also subject to other penalties
C['DEAD_TIME'] = 45
#
#
# PROJECT REALITY TK RESPAWN PENALTY
# Number of seconds added to the next respawn when teamkilling a player
# Default is 15 seconds
C['TEAMKILL_PENALTY'] = 15
#
#
# PROJECT REALITY SUICIDE RESPAWN PENALTY
# Number of seconds added to the next respawn when a player suicides
# Default is 30 seconds
C['SUICIDE_PENALTY'] = 15
#
#
# PROJECT REALITY REINFORCEMENTS RESPAWN PENALTY
# Edit these values to add seconds to player spawn times
# Define as 0 (zero) to disable
# Defaults
# 3 seconds per death
# -3 seconds per cp capture or destroyed objective
# -1 per defense action
# -10 for building
# 15 seconds maximum
C['SPAWN_PENALTY_DEATH'] = 3
C['SPAWN_PENALTY_CAPTURE'] = -3
C['SPAWN_PENALTY_OBJECTIVE'] = -3
C['SPAWN_PENALTY_DEFEND'] = -1
C['SPAWN_PENALTY_BUILD'] = -10
C['SPAWN_PENALTY_CAP'] = 15
#
#
# PROJECT REALITY MAX RESPAWN PENALTY
# Maximum number of seconds added to a respawn
# Default is 300 seconds
C['MAX_PENALTY'] = 300
#
#
# PROJECT REALITY REQUEST SPAM PENALTY
# Time in seconds (SPAM_PENALTY) the player is blocked becaues of a limited number of requests (SPAM_LIMIT) in a defined interval (SPAM_INTERVAL)
# Defaults:
# SPAM_LIMIT = 5
# SPAM_INTERVAL = 8
# SPAM_PENALTY = 30
C['SPAM_LIMIT'] = 5
C['SPAM_INTERVAL'] = 8
C['SPAM_PENALTY'] = 30
#
#
# PROJECT REALITY HEALTH UPON REVIVE SETTINGS
# Edit this setting to determine the health of players when they are revived
# Default is 10.1 health
C['REVIVE_HEALTH'] = 10.1
#
#
# PROJECT REALITY REVIVE TIME
# Time in seconds the player needs to stay without getting shot after revive or else he gets killed instantly.
# Default is 120 seconds
C['REVIVE_TIME'] = 120
#
#
# PROJECT REALITY GENERAL TICKETS
# Tickets lost when destroyed/killed
#
# Defaults:
# -10 = tanks, jets, attack helicopters, ifvs
# -5 = apcs, aavs, transport helicopters, recon
# -2 = jeeps, trucks
# -1 = soldiers
C['TICKETS'] = {
    CONSTANTS.VEHICLE_TYPE_ARMOR: -10,
    CONSTANTS.VEHICLE_TYPE_IFV: -10,
    CONSTANTS.VEHICLE_TYPE_JET: -10,
    CONSTANTS.VEHICLE_TYPE_HELIATTACK: -10,
    CONSTANTS.VEHICLE_TYPE_TURBOPROP: -7,
    CONSTANTS.VEHICLE_TYPE_SHIP: -50,
    CONSTANTS.VEHICLE_TYPE_HELI: -5,
    CONSTANTS.VEHICLE_TYPE_RECON: -5,
    CONSTANTS.VEHICLE_TYPE_AAV: -5,
    CONSTANTS.VEHICLE_TYPE_APC: -5,
    CONSTANTS.VEHICLE_TYPE_AFV: -2,
    CONSTANTS.VEHICLE_TYPE_ALC: -2,
    CONSTANTS.VEHICLE_TYPE_TRANSPORT: -2,
    CONSTANTS.VEHICLE_TYPE_SOLDIER: -1,
    CONSTANTS.VEHICLE_TYPE_FREE: 0

}
#
#
# Tickets lost when losing a flag (both neutralized and captured)
# Default is 30 tickets
C['TICKETS_CP'] = 30
#
#
# PROJECT REALITY SCORING SETTINGS
# Enable/Disable general score
# Default is enabled
C['SCORING_GENERAL'] = 1
#
#
# Enable/Disable teamwork score
# Default is enabled
C['SCORING_TEAMWORK'] = 1
#
#
# Enable/Disable kill count
# Default is enabled
C['SCORING_KILLS'] = 1
#
#
# Enable/Disable death count
# Default is enabled
C['SCORING_DEATHS'] = 1
#
#
# PROJECT REALITY ASSAULT AND SECURE COMMANDER
# Enable/Disable commander in AAS
# Default is enabled
C['AAS_COMMANDER'] = 1
#
#
# PROJECT REALITY INSURGENCY COMMANDER
# Enable/Disable commander in Insurgency
# Default is enabled
C['INSURGENCY_COMMANDER'] = 1
#
#
# PROJECT REALITY COUNTER-ATTACK COMMANDER
# Enable/Disable commander in Counter-Attack
# Default is enabled
C['COUNTER_COMMANDER'] = 1
#
#
# PROJECT REALITY COOP COMMANDER
# Enable/Disable commander in Coop
# Default is enabled
C['COOP_COMMANDER'] = 1
#
#
# PROJECT REALITY SCENARIO COMMANDER
# Enable/Disable commander in Scenario
# Default is enabled
C['SCENARIO_COMMANDER'] = 1
#
#
# PROJECT REALITY SKIRMISH COMMANDER
# Enable/Disable commander in Skirmish
# Default is enabled
C['SKIRMISH_COMMANDER'] = 1
#
#
# PROJECT REALITY CNC COMMANDER
# Enable/Disable commander in Command and Control
# Default is enabled
C['CNC_COMMANDER'] = 1
#
#
# PROJECT REALITY VEHICLE WARFARE COMMANDER
# Enable/Disable commander in Vehicle Warfare
# Default is enabled
C['VEHICLES_COMMANDER'] = 1
#
#
# PROJECT REALITY OBJECTIVE COMMANDER
# Enable/Disable commander in Objective
# Default is enabled
C['OBJECTIVE_COMMANDER'] = 1
#
#
# PROJECT REALITY ASSAULT AND SECURE SETTINGS
#
# Capture multiplier. The larger, the faster a flag can be capped at max speed
# Default is 1.5
C['AAS_MAX_CAPTURE_MULTIPLIER'] = 1.5
#
#
# Neutralise multiplier. The larger, the faster a flag can be capped at max speed
# Default is 1.5
C['AAS_MAX_NEUTRAL_MULTIPLIER'] = 1.5
#
#
# Number of players needed to capture or neutralize a CP
# Default is 2
C['AAS_MINNRTOTAKECONTROL'] = 2
#
#
# PROJECT REALITY SKIRMISH SETTINGS
#
# Number of tickets for each team
# Default is 150
C['SKIRMISH_TICKETS'] = 150
#
#
# Tickets lost when losing a flag (both neutralized and captured)
# Default is 10 tickets
C['SKIRMISH_TICKETS_CP'] = 10
#
#
# PROJECT REALITY VEHICLE WARFARE SETTINGS
#
# Number of tickets for each team
# Default is 200
C['VEHICLES_TICKETS'] = 200
#
#
# Time to neutralize a control point
# Default is 15
C['VEHICLES_NEUTRALIZE'] = 15
#
#
# Time to capture a control point
# Default is 15
C['VEHICLES_CAPTURE'] = 15
#
#
# Minimum number of players to take control of a flag
# Default is 6
C['VEHICLES_MINNRTOTAKECONTROL'] = 6
#
#
# Only kits allowed to be requested
# Default is tanker, pilot and officer
C['VEHICLES_KITS'] = ['tanker', 'pilot', 'officer']
#
#
# PROJECT REALITY INSURGENCY SETTINGS
#
# Number of objectives to destroy.
# Default is 5
C['INSURGENCY_OBJECTIVES'] = 5
#
#
# Max number of caches active at the same time
# Default is 2
C['INSURGENCY_ACTIVE'] = 2
#
#
# Number of points taken from the insurgent that destroys an objective
# Default is -100 points
C['INSURGENCY_TREASON_POINTS'] = -100
#
#
# Number of seconds added to the insurgent that destroys an objective
# Default is 300 seconds
C['INSURGENCY_TREASON_PENALTY'] = 300
#
#
# Number of points to the player when destroying an objective
# Default is 150 points
C['INSURGENCY_DESTROY_POINTS'] = 150
#
#
# Number of tickets added to team 2 when an objective is destroyed
# Default is 30
C['INSURGENCY_DESTROY_TICKETS'] = 30
#
#
# How much intel needed to reveal an objective on the map
# Default is 50 intel points
C['INSURGENCY_REVEAL_INTEL'] = 50
#
#
# How many seconds it takes for the attacking team to receive the information about a revealed objective
# Default is 300 seconds
C['INSURGENCY_REVEAL_INTERVAL'] = 300
#
#
# How much intel for capturing an insurgent
# Default is 10 intel points
C['INSURGENCY_INTEL_CAPTURE'] = 10
#
#
# How much intel lost for killing a civilian
# Default is -5 intel points
C['INSURGENCY_INTEL_CIVILIAN'] = -5
#
#
# How much intel for killing a normal insurgent
# Default is 1 intel point
C['INSURGENCY_INTEL_KILL'] = 1
#
#
# Which vehicle types are excluded from generating intel
# Default are air assets, tanks, ifvs, apcs and at vehicles
C['INSURGENCY_INTEL_EXCLUDE_VEHICLES'] = ['jet', 'ahe', 'the', 'tnk', 'atm']
#
#
# Maximum distance at which you gain intel from kills
# Default is 300m
C['INSURGENCY_INTEL_MAX_RANGE'] = 1500
#
#
# Which vehicle types are considered civilians
# Default are all unarmed ones
C['INSURGENCY_CIV_CARS'] = ['civ_bik_dirtbike', 'civ_bik_atv', 'civ_bik_hondacb500', 'civ_jep_forklift',
                            'civ_jep_support', 'civ_jep_car_white', 'civ_jep_car_blue', 'civ_jep_car_black',
                            'civ_jep_car', 'civ_trk_dumpster', 'civ_trk_semi',
                            'civ_jep_zastava900ak']
#
#
# How much intel lost for killing a civilian car
# Default is -5 intel points
C['INSURGENCY_INTEL_CIV_CARS'] = -5
#
#
# How many seconds it takes for a civilian car to be outside ROE
# Default is 60 seconds
C['INSURGENCY_INTERVAL_CIV_CARS'] = 60
#
#
# Points lost for killing a civilian car
# Default is 50m
C['INSURGENCY_DESTROY_CIV_CARS'] = -50
#
#
# PROJECT REALITY COMMAND AND CONTROL SETTINGS
#
# Maximum number of forward outposts
# Default is 1
C['CNC_OUTPOSTS_MAX'] = 1
#
#
# Minimum number of forward outposts needed to start the bleed
# Default is 0
C['CNC_OUTPOSTS_MIN'] = 0
#
#
# How many tickets lost per second when bleeding
# Default is -0.2
C['CNC_BLEED_TICKETS'] = -0.2
#
#
# How many tickets lost when outpost is destroyed
# Default is 50
C['CNC_OUTPOST_TICKETS'] = 50
#
#
# The multiplier for the number of defenses allowed per outpost
# Default is 2
C['CNC_DEFENSES_MULTIPLIER'] = 2
#
#
# Minimum distance an asset can be deployed from an outpost
# Default is DISTANCE_WIDE * 2
C['CNC_OUTPOST_DISTANCE'] = CONSTANTS.DISTANCE_WIDE * 2
#
#
# Minimum distance an asset can be deployed from a supply depot
# Default is DISTANCE_CLOSE
C['CNC_DEPOT_DISTANCE'] = CONSTANTS.DISTANCE_CLOSE
#
#
# Minimum distance an asset can be deployed from a command post
# Default is DISTANCE_CLOSE
C['CNC_COMMANDPOST_DISTANCE'] = CONSTANTS.DISTANCE_CLOSE
#
#
# Minimum distance an asset can be deployed from the map edge
# Default is DISTANCE_WIDE
C['CNC_EDGE_DISTANCE'] = CONSTANTS.DISTANCE_WIDE
#
#
# Interval of time to start checking for bleed
# Default is 600
C['CNC_START'] = 600
#
#
# Interval of time the team has to rebuild the FO before it starts bleeding
# Default is 300
C['CNC_DESTROYED_INTERVAL'] = 300
#
#
# PROJECT REALITY KIT REQUEST SYSTEM SETTINGS
#
#
# Team names that can't request kits
# Default is meinsurgent
C['KIT_REQUEST_BLOCK'] = []
#
#
# These are the maximum numbers of each of the limited kits that are available.
# These are done on a per team basis, with the number of players rounded up to 8, 16, 24, or 32.
# So if there are 20 players on the team, the limits in KIT_LIMIT_24 will be used.
# Default for 8 players: 0 = special kits, 1 = infantry kits
# Default for 16 players: 1 = special kits, 2 = infantry kits
# Default for 24 players: 1 = special kits, 3 = infantry kits
# Default for 32 players: 2 = special kits, 3 = infantry kits
# Default for 44 players: 2 = special kits, 4 = infantry kits
C['KIT_LIMIT_8'] = { 'sniper':     0, 'aa': 0, 'at': 0, 'engineer': 0, 'marksman': 1, 'assault': 1, 'riflemanat': 1,
                     'riflemanap': 1, 'mg': 1, 'spotter': 1 }
C['KIT_LIMIT_16'] = { 'sniper':     1, 'aa': 1, 'at': 1, 'engineer': 1, 'marksman': 2, 'assault': 2, 'riflemanat': 2,
                      'riflemanap': 2, 'mg': 2, 'spotter': 1 }
C['KIT_LIMIT_24'] = { 'sniper':     1, 'aa': 1, 'at': 1, 'engineer': 1, 'marksman': 3, 'assault': 4, 'riflemanat': 4,
                      'riflemanap': 3, 'mg': 3, 'spotter': 1 }
C['KIT_LIMIT_32'] = { 'sniper':     2, 'aa': 2, 'at': 1, 'engineer': 2, 'marksman': 3, 'assault': 5, 'riflemanat': 5,
                      'riflemanap': 3, 'mg': 3, 'spotter': 2 }
C['KIT_LIMIT_44'] = { 'sniper':     2, 'aa': 2, 'at': 1, 'engineer': 2, 'marksman': 4, 'assault': 8, 'riflemanat': 8,
                      'riflemanap': 4, 'mg': 4, 'spotter': 2 }
#
#
# The limited kits available for each faction and the number of players in the squad to be able to request it
# Defaults:
# 4 = infantry kits
# 2 = special kits
# 1 = vehicle kits
C['KIT_LIMITS'] = {
    'ch':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'gb':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'ger':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'mec':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'us':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'usa':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'arf':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'rifleman': 0
    },
    'cf':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'pl':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'ru':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'idf':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'chinsurgent': {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'rifleman': 0
    },
    'hamas':       {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'rifleman': 0
    },
    'taliban':     {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'rifleman': 0
    },
    'meinsurgent': {
        'officer': 2, 'insurgent1': 0, 'insurgent2': 0, 'insurgent3': 0, 'insurgent4': 0, 'sapper': 0,
        'medic':   2
    },
    'vnusa':       {
        'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6, 'spotter': 3,
        'officer': 2, 'sniper': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'vnusmc':      {
        'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6, 'spotter': 3,
        'officer': 2, 'sniper': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'vnnva':       {
        'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6, 'spotter': 3,
        'officer': 2, 'sniper': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'gb82':        {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'arg82':       {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'fr':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'fsa':         {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'rifleman': 0
    },
    'nl':          {
        'marksman': 6, 'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6,
        'spotter':  3,
        'officer':  2, 'sniper': 3, 'aa': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'ww2ger':      {
        'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6, 'spotter': 3,
        'officer': 2, 'sniper': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
    'ww2usa':      {
        'support': 4, 'riflemanat': 4, 'assault': 6, 'riflemanap': 4, 'specialist': 4, 'mg': 6, 'spotter': 3,
        'officer': 2, 'sniper': 3, 'at': 3, 'engineer': 3, 'medic': 2, 'tanker': 1, 'pilot': 1, 'rifleman': 0
    },
}
#
#
# The max number of each limited kits available for each squad
# Default is 1
C['KIT_LIMITS_SQUAD'] = {
    'sniper':   2, 'aa': 1, 'at': 1, 'engineer': 2, 'spotter': 1,
    'marksman': 1, 'riflemanat': 1, 'assault': 1, 'support': 1, 'medic': 2, 'riflemanap': 1, 'specialist': 1, 'mg': 1
}
#
#
#
# Per map kit limit override
C['KIT_LIMITS_MAPOVERRIDE'] = {
    # example
    #'muttrah_city_2': {
    #    'gpm_cq_64_team_1': {'sniper': 8}
    #    'gpm_cq_64_team_1': {'sniper': 8}
    #}
    'test_bootcamp': {
        'gpm_cq_64_team_1': {'sniper': 8},
        'gpm_cq_64_team_2': {'sniper': 8}
    }
}
#
#
# The amount of seconds a kit is available again for request
# Default is 600 seconds for special kits, 300 seconds for infantry kits
C['KIT_ALLOCATION_DELAY'] = {
    'sniper':     600, 'aa': 600, 'at': 600, 'engineer': 600, 'mg': 600, 'spotter': 600, 'marksman': 300, 'riflemanap': 300,
    'riflemanat': 300, 'assault': 300, 'support': 300, 'specialist': 300
}
#
#
# The amount of seconds a kit is available again for pickup
# Default is 300 seconds for special kits, 30 seconds for infantry kits
C['KIT_PICKUP_DELAY'] = {
    'sniper':     300, 'aa': 300, 'at': 300, 'engineer': 300, 'mg': 300, 'spotter': 300,
    'officer':    30, 'marksman': 30, 'riflemanat': 30, 'riflemanap': 30, 'assault': 30, 'medic': 30, 'support': 30,
    'specialist': 30,
    'tanker':     30, 'pilot': 30, 'unarmed': 30
}
# The amount of seconds the player needs to wait to request a kit after joining a squad
# Default is 120 seconds
C['KIT_SQUAD_DELAY'] = 120
#
# Define if a player can only pick up a kit from his faction
# Default is enabled
C['KIT_FACTION_LOCKED'] = 1
#
#
# These are the other objects from which kits can be requested. Following each name is the
# the maximum distance from the object in metres. Again, names should be lowercase.
# Default is DISTANCE_PICKUP or DISTANCE_SPAWN
C['KIT_SUPPLY_OBJECTS'] = {
    'us':          { 'pr_supply_crate_us':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_us': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_us': CONSTANTS.DISTANCE_PICKUP },
    'usa':         { 'pr_supply_crate_us':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_us': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_us': CONSTANTS.DISTANCE_PICKUP },
    'cf':          { 'pr_supply_crate_cf':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_cf': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_cf': CONSTANTS.DISTANCE_PICKUP },
    'gb':          { 'pr_supply_crate_gb':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_gb': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_gb': CONSTANTS.DISTANCE_PICKUP },
    'ger':         { 'pr_supply_crate_ger':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_ger': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_ger': CONSTANTS.DISTANCE_PICKUP },
    'ch':          { 'pr_supply_crate_ch':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_ch': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_ch': CONSTANTS.DISTANCE_PICKUP },
    'mec':         { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_mec': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_mec': CONSTANTS.DISTANCE_PICKUP, 'fixed_supply_crate_mec': CONSTANTS.DISTANCE_PICKUP },
    'pl':          { 'pr_supply_crate_pl':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_pl': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_pl': CONSTANTS.DISTANCE_PICKUP },
    'ru':          { 'pr_supply_crate_ru':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_ru': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_ru': CONSTANTS.DISTANCE_PICKUP },
    'idf':         { 'pr_supply_crate_idf':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_idf': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_idf': CONSTANTS.DISTANCE_PICKUP },
    'arf':         { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP, 'fixed_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP },
    'chinsurgent': { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_mil': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_mil': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_mil': CONSTANTS.DISTANCE_PICKUP },
    'hamas':       { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP, 'fixed_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP },
    'taliban':     { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_tal': CONSTANTS.DISTANCE_PICKUP },
    'meinsurgent': { 'ammocache': CONSTANTS.DISTANCE_PICKUP },
    'vnusa':       { 'pr_supply_crate_us':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_us': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_us': CONSTANTS.DISTANCE_PICKUP },
    'vnusmc':      { 'pr_supply_crate_us':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_us': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_us': CONSTANTS.DISTANCE_PICKUP },
    'vnnva':       { 'ammocache':                CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_vnnva': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_vnnva': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_vnnva': CONSTANTS.DISTANCE_PICKUP },
    'arg82':       { 'pr_supply_crate_arg':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_arg': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_arg': CONSTANTS.DISTANCE_PICKUP },
    'gb82':        { 'pr_supply_crate_gb':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_gb': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_gb': CONSTANTS.DISTANCE_PICKUP },
    'fr':          { 'pr_supply_crate_fr':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_fr': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_fr': CONSTANTS.DISTANCE_PICKUP },
    'fsa':         { 'ammocache':              CONSTANTS.DISTANCE_PICKUP, 'pr_supply_crate_fsa': CONSTANTS.DISTANCE_PICKUP,
                     'light_supply_crate_fsa': CONSTANTS.DISTANCE_PICKUP, 'fixed_supply_crate_fsa': CONSTANTS.DISTANCE_PICKUP },
    'nl':          { 'pr_supply_crate_nl':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_nl': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_nl': CONSTANTS.DISTANCE_PICKUP },
    'ww2ger':      { 'pr_supply_crate_ww2ger':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_ww2ger': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_ww2ger': CONSTANTS.DISTANCE_PICKUP },
    'ww2usa':      { 'pr_supply_crate_ww2usa':    CONSTANTS.DISTANCE_PICKUP, 'light_supply_crate_ww2usa': CONSTANTS.DISTANCE_PICKUP,
                     'fixed_supply_crate_ww2usa': CONSTANTS.DISTANCE_PICKUP },
}
#
#
# Kit supply objects that are vehicles, so the distance check must be in relation to the rear of the vehicle
# Default is DISTANCE_PICKUP
C['KIT_SUPPLY_OBJECTS_VEHICLES'] = {
    'us':          { 'us_apc_lav25':   CONSTANTS.DISTANCE_PICKUP, 'us_ifv_m2a2': CONSTANTS.DISTANCE_PICKUP,
                     'us_apc_aavp7a1': CONSTANTS.DISTANCE_PICKUP },
    'usa':         { 'us_apc_stryker': CONSTANTS.DISTANCE_PICKUP, 'us_apc_stryker_mk19': CONSTANTS.DISTANCE_PICKUP,
                     'us_ifv_m2a2':    CONSTANTS.DISTANCE_PICKUP },
    'cf':          { 'cf_apc_lav3': CONSTANTS.DISTANCE_PICKUP },
    'gb':          { 'gb_apc_warrior': CONSTANTS.DISTANCE_PICKUP, 'gb_apc_warrior_cage': CONSTANTS.DISTANCE_PICKUP },
    'ger':         { 'ger_apc_fuchs': CONSTANTS.DISTANCE_PICKUP, 'ger_apc_puma': CONSTANTS.DISTANCE_PICKUP,
                     'ger_ifv_puma':  CONSTANTS.DISTANCE_PICKUP },
    'ch':          { 'ch_apc_wz551':  CONSTANTS.DISTANCE_PICKUP, 'ch_apc_wz551a': CONSTANTS.DISTANCE_PICKUP,
                     'ch_ifv_wz551b': CONSTANTS.DISTANCE_PICKUP, 'ch_ifv_type86': CONSTANTS.DISTANCE_PICKUP, 'ch_ifv_zbl08': CONSTANTS.DISTANCE_PICKUP },
    'mec':         { 'mec_apc_boragh':    CONSTANTS.DISTANCE_PICKUP, 'mec_apc_btr60': CONSTANTS.DISTANCE_PICKUP,
                     'mec_ifv_bmp3':      CONSTANTS.DISTANCE_PICKUP, 'mec_apc_mtlb': CONSTANTS.DISTANCE_PICKUP,
                     'mec_apc_mtlb_hmg':  CONSTANTS.DISTANCE_PICKUP, 'mec_apc_mtlb_30mm': CONSTANTS.DISTANCE_PICKUP,
                     'mec_ifv_bmp2m':     CONSTANTS.DISTANCE_PICKUP, 'mec_apc_bmp2': CONSTANTS.DISTANCE_PICKUP,
                     'mil_ifv_bmp1':      CONSTANTS.DISTANCE_PICKUP, 'mec_apc_btr80': CONSTANTS.DISTANCE_PICKUP,
                     'mec_apc_btr80_alt': CONSTANTS.DISTANCE_PICKUP, 'mec_apc_btr80a': CONSTANTS.DISTANCE_PICKUP, 
                     'mec_apc_mtplb': CONSTANTS.DISTANCE_PICKUP  },
    'pl':          { 'pl_apc_rosomak': CONSTANTS.DISTANCE_PICKUP,
                     'pl_ifv_bwp1':    CONSTANTS.DISTANCE_PICKUP },
    'ru':          { 'ru_apc_btr60':     CONSTANTS.DISTANCE_PICKUP, 'ru_apc_btr80': CONSTANTS.DISTANCE_PICKUP,
                     'ru_apc_btr80_alt': CONSTANTS.DISTANCE_PICKUP, 'ru_apc_btr80a': CONSTANTS.DISTANCE_PICKUP,
                     'ru_apc_btr82am': CONSTANTS.DISTANCE_PICKUP,   'ru_ifv_bmp3m': CONSTANTS.DISTANCE_PICKUP,
                     'ru_ifv_bmp3':      CONSTANTS.DISTANCE_PICKUP, 'ru_apc_mtlb': CONSTANTS.DISTANCE_PICKUP,
                     'ru_apc_mtlb_hmg':  CONSTANTS.DISTANCE_PICKUP, 'ru_ifv_bmp2': CONSTANTS.DISTANCE_PICKUP,
                     'ru_apc_bmp2':      CONSTANTS.DISTANCE_PICKUP, 'ru_apc_mtplb': CONSTANTS.DISTANCE_PICKUP,
                     'ru_ifv_bmp2m':      CONSTANTS.DISTANCE_PICKUP },
    'idf':         { 'idf_apc_m113':  CONSTANTS.DISTANCE_PICKUP, 'idf_apc_m113_logistics': CONSTANTS.DISTANCE_PICKUP,
                     'idf_apc_namer': CONSTANTS.DISTANCE_PICKUP },
    'arf':         { },
    'chinsurgent': { 'mil_ifv_bmp1': CONSTANTS.DISTANCE_PICKUP },
    'hamas':       { },
    'fsa':         { 'fsa_ifv_bmp1': CONSTANTS.DISTANCE_PICKUP },
    'taliban':     { },
    'meinsurgent': { },
    'vnusa':       { 'us_apc_m113': CONSTANTS.DISTANCE_PICKUP, 'us_apc_acav': CONSTANTS.DISTANCE_PICKUP },
    'vnusmc':      { 'us_apc_m113': CONSTANTS.DISTANCE_PICKUP, 'us_apc_acav': CONSTANTS.DISTANCE_PICKUP },
    'vnnva':       { 'nva_apc_btr60': CONSTANTS.DISTANCE_PICKUP },
    'gb82':        { },
    'arg82':       { 'arg_apc_lvtp7': CONSTANTS.DISTANCE_PICKUP },
    'fr':          { 'fr_apc_vab': CONSTANTS.DISTANCE_PICKUP, 'fr_apc_vbci': CONSTANTS.DISTANCE_PICKUP },
    'nl':          { 'nl_apc_boxer':      CONSTANTS.DISTANCE_PICKUP, 'nl_apc_boxer_unarmed': CONSTANTS.DISTANCE_PICKUP,
                     'nl_apc_cv90':       CONSTANTS.DISTANCE_PICKUP, 'nl_apc_ypr50': CONSTANTS.DISTANCE_PICKUP,
                     'nl_apc_ypr50_gpmg': CONSTANTS.DISTANCE_PICKUP },
    'ww2ger':      { 'ger_apc_251c': CONSTANTS.DISTANCE_PICKUP, 'ger_apc_251c_alt': CONSTANTS.DISTANCE_PICKUP },
    'ww2usa':      { 'us_apc_m3': CONSTANTS.DISTANCE_PICKUP },
}
#
#
# Vehicles that have side doors instead of back doors.
# Default is ru_apc_btr60, ru_apc_btr80, ru_apc_btr80_alt, ru_apc_btr80a, mec_apc_btr60
C['KIT_SUPPLY_OBJECTS_VEHICLES_SIDEDOORS'] = ['ru_apc_btr60', 'ru_apc_btr80', 'ru_apc_btr80_alt', 'ru_apc_btr80a',
                                              'mec_apc_btr60', 'mec_apc_btr80']
#
#
# Kit supply objects that don't check for the object team ownership
# Default is DISTANCE_PICKUP or DISTANCE_SPAWN
C['KIT_SUPPLY_OBJECTS_OPEN'] = {
    'us':          { 'vehicle_depot_us': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'usa':         { 'vehicle_depot_us': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'cf':          { 'vehicle_depot_cf': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'gb':          { 'vehicle_depot_gb': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'ger':         { 'vehicle_depot_ger': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'ch':          { 'vehicle_depot_ch': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'mec':         { 'vehicle_depot_mec': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'pl':          { 'vehicle_depot_pl': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'ru':          { 'vehicle_depot_ru': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'idf':         { 'vehicle_depot_idf': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'arf':         { 'vehicle_depot_tal': CONSTANTS.DISTANCE_SPAWN },
    'chinsurgent': { 'vehicle_depot_mil': CONSTANTS.DISTANCE_SPAWN },
    'hamas':       { 'vehicle_depot_tal': CONSTANTS.DISTANCE_SPAWN },
    'taliban':     { 'vehicle_depot_tal': CONSTANTS.DISTANCE_SPAWN },
    'meinsurgent': { },
    'vnusa':       { 'vehicle_depot_us': CONSTANTS.DISTANCE_SPAWN },
    'vnusmc':      { 'vehicle_depot_us': CONSTANTS.DISTANCE_SPAWN },
    'vnnva':       { 'vehicle_depot_vnnva': CONSTANTS.DISTANCE_SPAWN },
    'gb82':        { 'vehicle_depot_gb': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'arg82':       { 'vehicle_depot_arg': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'fr':          { 'vehicle_depot_fr': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'fsa':         { 'vehicle_depot_fsa': CONSTANTS.DISTANCE_SPAWN },
    'nl':          { 'vehicle_depot_nl': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'ww2ger':      { 'vehicle_depot_ww2ger': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
    'ww2usa':      { 'vehicle_depot_ww2usa': CONSTANTS.DISTANCE_SPAWN, 'para_supply_crate': CONSTANTS.DISTANCE_PICKUP },
}
#
#
# Min time the player needs to wait between spawning kits in seconds.
# Default is 120 seconds (2 minutes)
C['KIT_REQUEST_INTERVAL'] = 120
#
#
# PROJECT REALITY VEHICLES SYSTEMS SETTINGS
#
# Enable/Disable the vehicles kits, squad and roles requirements.
# Default is enabled
C['VEHICLES_REQUIREMENTS'] = 1
#
#
# Enable/Disable the vehicles damage system
# Default is enabled
C['VEHICLES_DAMAGE'] = 1
#
#
# Enable/Disable the check for damage on exiting a vehicle at high speed
# Default is enabled
C['VEHICLES_EXIT_DAMAGE_SPEED'] = 1
#
#
# Enable/Disable the check for damage on exiting a vehicle that is critical damaged
# Default is enabled
C['VEHICLES_EXIT_DAMAGE_CRITICAL'] = 1
#
#
# Enable/Disable start delay on vehicles
# Default is enabled
C['VEHICLES_START_DELAY'] = 1
#
#
# PROJECT REALITY RALLY POINT SYSTEM SETTINGS
#
# Teams that use the squad rally point system
# Default is arf, ch, gb, mec, us, usa, cf, chinsurgent, ru, arf, taliban, idf, hamas, ger, vnusa, vnusmc, vnnva, gb82, arg82, fr, nl, ww2ger, ww2usa
C['RALLY_TEAMS'] = ['arf', 'arg82', 'cf', 'ch', 'chinsurgent', 'fr', 'fsa', 'gb', 'gb82', 'ger', 'hamas', 'idf', 'mec',
                    'nl', 'pl', 'ru', 'taliban', 'us', 'usa', 'vnnva', 'vnusa', 'vnusmc', 'ww2ger', 'ww2usa']
#
#
# Teams that use the commander rally point system
# Default is arf, ch, gb, mec, us, usa, cf, chinsurgent, ru, arf, taliban, idf, hamas, ger, vnusa, vnusmc, vnnva, gb82, arg82, fr, nl, ww2ger, ww2usa
C['RALLY_TEAMS_COMMANDER'] = ['arf', 'arg82', 'cf', 'ch', 'chinsurgent', 'fr', 'fsa', 'gb', 'gb82', 'ger', 'hamas',
                              'idf', 'mec', 'nl', 'pl', 'ru', 'taliban', 'us', 'usa', 'vnnva', 'vnusa',
                              'vnusmc', 'ww2ger', 'ww2usa']

#
# Min number of players required in the squad to set a squad rally point
# Default is 2
C['RALLY_LIMIT_SQUAD'] = 2
#
#
# Min number of players required in the team to set a commander rally point
# Default is 12
C['RALLY_LIMIT_COMMANDER'] = 12
#
#
# Min number of close by squad members when setting rally points (must be lower than RALLY_LIMIT_SQUAD)
# Default is 2
C['RALLY_CLOSE_SQUAD'] = 2
#
#
# Min number of close by team members when setting rally points (must be equal or lower than 32)
# Default is 6
C['RALLY_CLOSE_COMMANDER'] = 6
#
#
# Min number of close by squad leaders when setting rally points (must be equal or lower than 9)
# Default is 1
C['RALLY_CLOSE_COMMANDER_SL'] = 1
#
#
# Number of soldiers needed to be close to a rally point to automatically delete it
# Default is 1
C['RALLY_CLOSE_DESTROY'] = 1
#
#
# Radius a rallypoint gets destroyed on >=4km maps
# Default is 125m
C['RALLY_AREA_HUGE'] = 125
#
#
# Radius a rallypoint gets destroyed on 2km maps
# Default is 125m
C['RALLY_AREA_BIG'] = 125
#
#
# Radius a rallypoint gets destroyed on <=1km maps
# Default is 50m
C['RALLY_AREA_SMALL'] = 50
#
#
# Radius a rallypoint is supported from on >=4km maps
# Default is 600m
C['RALLY_SUPPORT_AREA_HUGE'] = 600
#
#
# Radius a rallypoint is supported from on 2km maps
# Default is 300m
C['RALLY_SUPPORT_AREA_BIG'] = 300
#
#
# Radius a rallypoint is supported from on <=1km maps
# Default is 150m
C['RALLY_SUPPORT_AREA_SMALL'] = 150
#
#
# The multiplier on how far a rally must be set from enemies, relative to RALLY_AREA_*
# Default is 1
C['RALLY_SET_MULTIPLIER'] = 1
#
#
# How many seconds between setting RPs
# Default is 60 seconds
C['RALLY_INTERVAL'] = 60
#
#
# How many seconds between setting RPs when overrun
# Default is 300 seconds
C['RALLY_OVERRUN_INTERVAL'] = 300
#
#
# How many seconds a squad rally expires after being set
# Default is 60 seconds
C['RALLY_EXPIRATION'] = 60
#
#
# How many seconds a commander rally expires after being set
# Default is 60 seconds
C['RALLY_EXPIRATION_COMMANDER'] = 60
#
#
# How many seconds expirable mapper rally points will stay in play after round start
# Default is 300 seconds (5 minutes)
C['RALLY_MAPPER_EXPIRATION'] = 300
#
#
# How many players can spawn on a limited mapper placed rally point
# Default is 12
C['RALLY_MAPPER_LIMITED'] = 12
#
#
# Number that divides the amount of random rally points in a map (can be set by mapper)
# Default is 3
C['RALLY_RANDOM'] = 3
#
#
# The teams that have pickup kits at random rally points, and the list of kits to be used.
# Default is {}
C['RALLY_RANDOM_PICKUP'] = { }
#
#
# How many seconds expirable paradrop spawn points will stay in play after round start
# Default is 120 seconds (2 minutes)
C['PARADROP_EXPIRATION'] = 120
#
#
# What vehicle types support rally points
# Default is apc, ifv
C['RALLY_SUPPORT_VEHICLE_TYPES'] = ['apc', 'ifv']
#
#
# PROJECT REALITY COMMANDER ASSETS SYSTEM SETTINGS
#
# The teams that use the commander assets system
# Default is ch, gb, mec, us, usa, cf, chinsurgent, meinsurgent, ru, arf, taliban, idf, hamas, ger, vnusa, vnusmc, vnnva, gb82, arg82, fr, nl, ww2ger, ww2usa
C['ASSET_TEAMS'] = ['ch', 'gb', 'mec', 'us', 'usa', 'fsa', 'cf', 'chinsurgent', 'meinsurgent', 'pl', 'ru', 'arf', 'taliban',
                    'idf', 'hamas', 'ger', 'vnusa', 'vnusmc', 'vnnva', 'gb82', 'arg82', 'fr', 'nl', 'ww2ger', 'ww2usa']
#
#
# The teams that do not require crates to build outposts
# Default is fsa
C['ASSET_TEAMS_FREE_OUTPOST'] = ['fsa']
#
#
# Assets types that give points to defenders
# Default is outpost
C['ASSET_POINTS_DEFEND'] = ['outpost']
#
#
# Assets types that give points to commander for building
# Default is outpost
C['ASSET_POINTS_BUILD'] = ['outpost']
#
#
# Minimum time before you gain points for building assets after your last one
# Default is 300 seconds
C['ASSET_POINTS_INTERVAL'] = 300
#
#
# Assets types that spawn vehicles
# Default is outpost
C['ASSET_VEHICLES_LIST'] = { }
#
#
# Interval in seconds to check for the existence of vehicle spawners in assets
# Default is 60 seconds
C['ASSET_VEHICLES_INTERVAL'] = 60
#
#
# Asset requests that require cmdr approval
# Default is []
C['ASSET_ORDER_APPROVAL'] = []
#
#
# Asset maximum of forward outposts in map
# Default is 6
C['ASSET_MAX_OUTPOSTS'] = 6
#
#
# Asset maximum of static defenses in area
# Default is 10
C['ASSET_MAX_STATIC_DEFENSES'] = 10
#
#
# Asset maximum of heavy defenses in area
# Default is 1
C['ASSET_MAX_HEAVY_DEFENSES'] = 1
#
#
# Asset maximum of medium defenses in area
# Default is 1
C['ASSET_MAX_MEDIUM_DEFENSES'] = 1
#
#
# Asset maximum of light defenses in area
# Default is 2
C['ASSET_MAX_LIGHT_DEFENSES'] = 2
#
#
# Asset maximum of static defenses in the map
# Default is 100
C['ASSET_MAX_STATIC_DEFENSES_MAP'] = 100
#
#
# Asset maximum of heavy defenses in the map
# Default is 3
C['ASSET_MAX_HEAVY_DEFENSES_MAP'] = 3
#
#
# Asset maximum of mortars in the map
# Default is 2
C['ASSET_MAX_MORTARS_MAP'] = 2
#
#
# Asset maximum of medium defenses in the map
# Default is 12
C['ASSET_MAX_MEDIUM_DEFENSES_MAP'] = 12
#
#
# Asset maximum of light defenses in the map
# Default is 24
C['ASSET_MAX_LIGHT_DEFENSES_MAP'] = 24
#
#
# Minimum distance an asset can be deployed from an outpost
# Default is DISTANCE_WIDE
C['ASSET_OUTPOST_DISTANCE'] = CONSTANTS.DISTANCE_WIDE
#
#
# Minimum distance an asset can be deployed from a supply depot
# Default is DISTANCE_CLOSE
C['ASSET_DEPOT_DISTANCE'] = CONSTANTS.DISTANCE_CLOSE
#
#
# Minimum distance an asset can be deployed from a command post
# Default is DISTANCE_WIDE
C['ASSET_COMMANDPOST_DISTANCE'] = CONSTANTS.DISTANCE_WIDE
#
#
# Minimum distance an asset can be deployed from the map edge
# Default is DISTANCE_SPAWN
C['ASSET_EDGE_DISTANCE'] = CONSTANTS.DISTANCE_SPAWN
#
#
# The teams that can use the UAV type 1 system
# Default is ch, gb, mec, us, usa, cf, ru, idf, ger, fr, nl
C['ASSET_TEAMS_UAV1'] = ['ch', 'gb', 'mec', 'pl', 'us', 'usa', 'cf', 'ru', 'idf', 'ger', 'fr', 'nl']
#
#
# The teams that can use the UAV type 2 system
# Default is chinsurgent
C['ASSET_TEAMS_UAV2'] = ['chinsurgent']
#
#
# Maps that don't have UAVs available
# Default is asad_khal
C['ASSET_MAPS_UAV_NONE'] = ['asad_khal', 'assault_on_mestia', 'korengal', 'kozelsk', 'dovre']
#
#
# Maps that specifically use 1Km UAVs
# Default is albasrah
C['ASSET_MAPS_UAV_1KM'] = []
#
#
# Maps that specifically use 2Km UAVs
# Default is None
C['ASSET_MAPS_UAV_2KM'] = []
#
#
# Maps that specifically use 4Km UAVs
# Default is None
C['ASSET_MAPS_UAV_4KM'] = []
#
#
# COMMAND POST SETTINGS
#
# The command post templates for each team
# Default is deployable_commandpost and acv variants
C['COMMANDPOST_TEMPLATES'] = {
    'us':          ['deployable_commandpost_us', 'us_acv_lavc2'],
    'usa':         ['deployable_commandpost_us', 'us_acv_stryker'],
    'cf':          ['deployable_commandpost_cf', 'cf_acv_lav3cpv'],
    'gb':          ['deployable_commandpost_gb', 'gb_acv_sultan'],
    'ger':         ['deployable_commandpost_ger', 'ger_acv_m557a2'],
    'ch':          ['deployable_commandpost_ch', 'ch_acv_wz551', 'ch_ship_type75_lpd_atc'],
    'mec':         ['deployable_commandpost_mec', 'mec_acv_btr60pu'],
    'pl':          ['deployable_commandpost_pl', 'pl_acv_star1466'],
    'ru':          ['deployable_commandpost_ru', 'ru_acv_btr60pu','ru_ship_andreev_lpd_atc'],
    'idf':         ['deployable_commandpost_idf', 'idf_acv_m557'],
    'arf':         ['deployable_commandpost_meins', 'arf_acv_technical'],
    'chinsurgent': ['deployable_commandpost_chins', 'mil_acv_technical'],
    'hamas':       ['deployable_commandpost_meins', 'deployable_insurgent_hideout'],
    'taliban':     ['deployable_commandpost_meins', 'deployable_insurgent_hideout'],
    'meinsurgent': ['deployable_commandpost_meins', 'deployable_insurgent_hideout'],
    'fsa':         ['deployable_commandpost_meins', 'deployable_insurgent_hideout'],
    'gb82':        ['deployable_commandpost_gb', 'gb_acv_sultan'],
    'arg82':       ['deployable_commandpost_arg', 'us_acv_lavc2'],
    'fr':          ['deployable_commandpost_mec', 'fr_acv_vab'],
    'nl':          ['deployable_commandpost_nl', 'nl_acv_boxer'],
    'ww2ger':      ['deployable_commandpost_ww2ger', 'ger_acv_tent'],
    'ww2usa':      ['deployable_commandpost_ww2usa', 'us_acv_tent'],
}
#
#
# VEHICLE SUPPLY DEPOT SETTINGS
#
# The vehicle supply depot templates for each team
# Default is vehicle_depot
C['VEHICLE_SUPPLY_DEPOT_TEMPLATES'] = {
    'us':          ['vehicle_depot_us'],
    'usa':         ['vehicle_depot_us'],
    'cf':          ['vehicle_depot_cf'],
    'gb':          ['vehicle_depot_gb'],
    'ger':         ['vehicle_depot_ger'],
    'ch':          ['vehicle_depot_ch'],
    'mec':         ['vehicle_depot_mec'],
    'pl':          ['vehicle_depot_pl'],
    'ru':          ['vehicle_depot_ru'],
    'idf':         ['vehicle_depot_idf'],
    'fsa':         ['vehicle_depot_fsa'],
    'arf':         ['vehicle_depot_tal'],
    'chinsurgent': ['vehicle_depot_mil'],
    'hamas':       ['vehicle_depot_tal'],
    'taliban':     ['vehicle_depot_tal'],
    'meinsurgent': ['vehicle_depot_tal'],
    'vnusa':       ['vehicle_depot_us'],
    'vnusmc':      ['vehicle_depot_us'],
    'vnnva':       ['vehicle_depot_vnnva'],
    'gb82':        ['vehicle_depot_gb'],
    'arg82':       ['vehicle_depot_arg'],
    'fr':          ['vehicle_depot_fr'],
    'nl':          ['vehicle_depot_nl'],
    'ww2ger':      ['vehicle_depot_ww2ger'],
    'ww2usa':      ['vehicle_depot_ww2usa'],
}
#
#
# FORWARD OUTPOST SETTINGS
#
# The outpost templates for each team
# Default is deployable_firebase, deployable_insurgent_hideout
C['OUTPOST_TEMPLATES'] = {
    'us':          ['deployable_firebase'],
    'usa':         ['deployable_firebase'],
    'cf':          ['deployable_firebase'],
    'gb':          ['deployable_firebase'],
    'ger':         ['deployable_firebase'],
    'ch':          ['deployable_firebase'],
    'mec':         ['deployable_firebase'],
    'pl':          ['deployable_firebase'],
    'ru':          ['deployable_firebase'],
    'idf':         ['deployable_firebase'],
    'arf':         ['deployable_insurgent_hideout'],
    'chinsurgent': ['deployable_firebase'],
    'hamas':       ['deployable_insurgent_hideout'],
    'taliban':     ['deployable_insurgent_hideout'],
    'meinsurgent': ['deployable_insurgent_hideout'],
    'fsa':         ['deployable_insurgent_hideout'],
    'vnusa':       ['deployable_firebase'],
    'vnusmc':      ['deployable_firebase'],
    'vnnva':       ['deployable_firebase'],
    'gb82':        ['deployable_firebase'],
    'arg82':       ['deployable_firebase'],
    'fr':          ['deployable_firebase'],
    'nl':          ['deployable_firebase'],
    'ww2ger':      ['deployable_firebase'],
    'ww2usa':      ['deployable_firebase'],
}
# Hideout ammo regen
C['HIDEOUT_AMMO_PER_MINUTE'] = 200
#
#
# The outpost dummy templates for each team
# Default is deployable_firebase_dummy, deployable_bunker_dummy and deployable_insurgent_hideout_dummy
C['OUTPOST_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_firebase_dummy'],
    'usa':         ['deployable_firebase_dummy'],
    'cf':          ['deployable_firebase_dummy'],
    'gb':          ['deployable_firebase_dummy'],
    'ger':         ['deployable_firebase_dummy'],
    'ch':          ['deployable_firebase_dummy'],
    'mec':         ['deployable_firebase_dummy'],
    'pl':          ['deployable_firebase_dummy'],
    'ru':          ['deployable_firebase_dummy'],
    'idf':         ['deployable_firebase_dummy'],
    'arf':         ['deployable_insurgent_hideout_dummy'],
    'chinsurgent': ['deployable_firebase_dummy'],
    'hamas':       ['deployable_insurgent_hideout_dummy'],
    'taliban':     ['deployable_insurgent_hideout_dummy'],
    'meinsurgent': ['deployable_insurgent_hideout_dummy'],
    'fsa':         ['deployable_insurgent_hideout_dummy'],
    'vnusa':       ['deployable_firebase_dummy'],
    'vnusmc':      ['deployable_firebase_dummy'],
    'vnnva':       ['deployable_firebase_dummy'],
    'gb82':        ['deployable_firebase_dummy'],
    'arg82':       ['deployable_firebase_dummy'],
    'fr':          ['deployable_firebase_dummy'],
    'nl':          ['deployable_firebase_dummy'],
    'ww2ger':      ['deployable_firebase_dummy'],
    'ww2usa':      ['deployable_firebase_dummy'],
}
#
#
# Asset templates that are hideouts
# Default is 'deployable_insurgent_hideout', 'deployable_insurgent_hideout_sp'
C['ASSET_HIDEOUT_TEMPLATES'] = ['deployable_insurgent_hideout', 'deployable_insurgent_hideout_sp']
#
#
# Minimum distance between hideouts
# Default is DISTANCE_WIDE
C['HIDEOUT_DISTANCE'] = CONSTANTS.DISTANCE_WIDE
#
#
# Asset maximum of hideouts in map
# Default is 6
C['ASSET_MAX_HIDEOUTS'] = 6
# Number of soldiers needed to be very close to an outpost to automatically disable the spawn point
# Default is 1
# C['OUTPOST_VERY_CLOSE_DISABLE'] = 1
#
#
# Number of soldiers needed to be close to an outpost to automatically disable the spawn point
# Default is 2
C['OUTPOST_CLOSE_DISABLE'] = 2
#
#
# Number of soldiers needed to be close to an outpost to automatically disable the spawn point
# Default is 5
C['OUTPOST_FAR_DISABLE'] = 4
#
#
# Number of soldiers needed to be close to an outpost to automatically disable the spawn point
# Default is 8
C['OUTPOST_VERY_FAR_DISABLE'] = 8
#
#
# Distance to be considered very close to outpost
# Default is 10
# C['OUTPOST_VERY_CLOSE_DISTANCE'] = 10
#
#
# Distance to be considered close to outpost
# Default is 50
C['OUTPOST_CLOSE_DISTANCE'] = 50
#
#
# Distance to be considered far to outpost
# Default is 100
C['OUTPOST_FAR_DISTANCE'] = 100
#
#
# Distance to be considered very far to outpost
# Default is 150
C['OUTPOST_VERY_FAR_DISTANCE'] = 150
#
#
# How many seconds after building an outpost it will become spawnable
# Default is 90 seconds
C['OUTPOST_LOST_INTERVAL'] = 90
#
#
# How many seconds after disabling an outpost the spawn is automatically reenabled (if no enemies are close)
# Default is 90 seconds
C['OUTPOST_LOST_INTERVAL_OVERRUN'] = 30
#
#
# Minimum distance between forward outposts
# Default is DISTANCE_WIDE
C['OUTPOST_DISTANCE'] = CONSTANTS.DISTANCE_WIDE
#
#
# SUPPLIES SETTINGS
#
# The supplies templates required to build the forward outposts.
# Default is pr_supply_crate
C['SUPPLIES_TEMPLATES'] = {
    'us':          [['pr_supply_crate_us', 1], ['light_supply_crate_us', 0.5]],
    'usa':         [['pr_supply_crate_us', 1], ['light_supply_crate_us', 0.5]],
    'cf':          [['pr_supply_crate_cf', 1], ['light_supply_crate_cf', 0.5]],
    'gb':          [['pr_supply_crate_gb', 1], ['light_supply_crate_gb', 0.5]],
    'ger':         [['pr_supply_crate_ger', 1], ['light_supply_crate_ger', 0.5]],
    'ch':          [['pr_supply_crate_ch', 1], ['light_supply_crate_ch', 0.5]],
    'mec':         [['pr_supply_crate_mec', 1], ['light_supply_crate_mec', 0.5]],
    'pl':          [['pr_supply_crate_pl', 1], ['light_supply_crate_pl', 0.5]],
    'ru':          [['pr_supply_crate_ru', 1], ['light_supply_crate_ru', 0.5]],
    'idf':         [['pr_supply_crate_idf', 1], ['light_supply_crate_idf', 0.5]],
    'chinsurgent': [['pr_supply_crate_mil', 1], ['light_supply_crate_mil', 0.5]],
    'fsa':         [['pr_supply_crate_fsa', 1], ['light_supply_crate_fsa', 0.5]],
    'vnusa':       [['pr_supply_crate_us', 1], ['light_supply_crate_us', 0.5]],
    'vnusmc':      [['pr_supply_crate_us', 1], ['light_supply_crate_us', 0.5]],
    'vnnva':       [['pr_supply_crate_vnnva', 1], ['light_supply_crate_vnnva', 0.5]],
    'gb82':        [['pr_supply_crate_gb', 1], ['light_supply_crate_gb', 0.5]],
    'arg82':       [['pr_supply_crate_arg', 1], ['light_supply_crate_arg', 0.5]],
    'fr':          [['pr_supply_crate_fr', 1], ['light_supply_crate_fr', 0.5]],
    'nl':          [['pr_supply_crate_nl', 1], ['light_supply_crate_nl', 0.5]],
    'ww2ger':      [['pr_supply_crate_ww2ger', 1], ['light_supply_crate_ww2ger', 0.5]],
    'ww2usa':      [['pr_supply_crate_ww2usa', 1], ['light_supply_crate_ww2usa', 0.5]],
}
#
#
# The supply objects templates that are team locked.
# Defaults are pr_supply_crate, pr_supply_depot, light_supply_crate and fixed_supply_crate
C['SUPPLIES_TEMPLATES_TEAMLOCKED'] = {
    'us':          ['pr_supply_crate_us', 'pr_supply_depot_us', 'light_supply_crate_us', 'fixed_supply_crate_us'],
    'usa':         ['pr_supply_crate_us', 'pr_supply_depot_us', 'light_supply_crate_us', 'fixed_supply_crate_us'],
    'cf':          ['pr_supply_crate_cf', 'pr_supply_depot_cf', 'light_supply_crate_cf', 'fixed_supply_crate_cf'],
    'gb':          ['pr_supply_crate_gb', 'pr_supply_depot_gb', 'light_supply_crate_gb', 'fixed_supply_crate_gb'],
    'ger':         ['pr_supply_crate_ger', 'pr_supply_depot_ger', 'light_supply_crate_ger', 'fixed_supply_crate_ger'],
    'ch':          ['pr_supply_crate_ch', 'pr_supply_depot_ch', 'light_supply_crate_ch', 'fixed_supply_crate_ch'],
    'mec':         ['pr_supply_crate_mec', 'pr_supply_depot_mec', 'light_supply_crate_mec', 'fixed_supply_crate_mec'],
    'pl':          ['pr_supply_crate_pl', 'pr_supply_depot_pl', 'light_supply_crate_pl', 'fixed_supply_crate_pl'],
    'ru':          ['pr_supply_crate_ru', 'pr_supply_depot_ru', 'light_supply_crate_ru', 'fixed_supply_crate_ru'],
    'idf':         ['pr_supply_crate_idf', 'pr_supply_depot_idf', 'light_supply_crate_idf', 'fixed_supply_crate_idf'],
    'chinsurgent': ['pr_supply_crate_mil', 'pr_supply_depot_mil', 'light_supply_crate_mil', 'fixed_supply_crate_mil'],
    'vnusa':       ['pr_supply_crate_us', 'pr_supply_depot_us', 'light_supply_crate_us', 'fixed_supply_crate_us'],
    'vnusmc':      ['pr_supply_crate_us', 'pr_supply_depot_us', 'light_supply_crate_us', 'fixed_supply_crate_us'],
    'vnnva':       ['pr_supply_crate_vnnva', 'pr_supply_depot_vnnva', 'light_supply_crate_vnnva', 'fixed_supply_crate_vnnva'],
    'taliban':     ['pr_supply_crate_tal', 'pr_supply_depot_tal', 'light_supply_crate_tal', 'fixed_supply_crate_tal'],
    'gb82':        ['pr_supply_crate_gb', 'pr_supply_depot_gb', 'light_supply_crate_gb', 'fixed_supply_crate_gb'],
    'arg82':       ['pr_supply_crate_arg', 'pr_supply_depot_arg', 'light_supply_crate_arg', 'fixed_supply_crate_arg'],
    'fr':          ['pr_supply_crate_fr', 'pr_supply_depot_fr', 'light_supply_crate_fr', 'fixed_supply_crate_fr'],
    'fsa':         ['pr_supply_crate_fsa', 'pr_supply_depot_fsa', 'light_supply_crate_fsa', 'fixed_supply_crate_fsa'],
    'nl':          ['pr_supply_crate_nl', 'pr_supply_depot_nl', 'light_supply_crate_nl', 'fixed_supply_crate_nl'],
    'ww2ger':      ['pr_supply_crate_ww2ger', 'pr_supply_depot_ww2ger', 'light_supply_crate_ww2ger', 'fixed_supply_crate_ww2ger'],
    'ww2usa':      ['pr_supply_crate_ww2usa', 'pr_supply_depot_ww2usa', 'light_supply_crate_ww2usa', 'fixed_supply_crate_ww2usa'],
}
#
#
# ANTI-AIR SETTINGS
#
# The anti-air templates for each team
# Default is deployable_stinger and deployable_djigit
C['ANTIAIR_TEMPLATES'] = {
    'us':          ['deployable_stinger'],
    'usa':         ['deployable_stinger'],
    'cf':          ['deployable_stinger'],
    'gb':          ['deployable_stinger'],
    'ger':         ['deployable_stinger'],
    'ch':          ['deployable_djigit'],
    'mec':         ['deployable_djigit'],
    'pl':          ['deployable_djigit'],
    'ru':          ['deployable_djigit'],
    'idf':         ['deployable_stinger'],
    'arf':         ['deployable_dshk'],
    'chinsurgent': ['deployable_dshk'],
    'hamas':       ['deployable_dshk'],
    'taliban':     ['deployable_dshk'],
    'meinsurgent': ['deployable_dshk'],
    'fsa':         ['deployable_dshk'],
    'vnnva':       ['deployable_dshk'],
    'gb82':        ['deployable_stinger'],
    'arg82':       ['deployable_djigit'],
    'fr':          ['deployable_mistral'],
    'nl':          ['deployable_stinger'],
}
#
#
# The anti-air dummy templates for each team
# Default is deployable_stinger_dummy and deployable_djigit_dummy
C['ANTIAIR_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_stinger_dummy'],
    'usa':         ['deployable_stinger_dummy'],
    'cf':          ['deployable_stinger_dummy'],
    'gb':          ['deployable_stinger_dummy'],
    'ger':         ['deployable_stinger_dummy'],
    'ch':          ['deployable_djigit_dummy'],
    'mec':         ['deployable_djigit_dummy'],
    'pl':          ['deployable_djigit_dummy'],
    'ru':          ['deployable_djigit_dummy'],
    'idf':         ['deployable_stinger_dummy'],
    'arf':         ['deployable_dshk_dummy'],
    'chinsurgent': ['deployable_dshk_dummy'],
    'hamas':       ['deployable_dshk_dummy'],
    'taliban':     ['deployable_dshk_dummy'],
    'meinsurgent': ['deployable_dshk_dummy'],
    'fsa':         ['deployable_dshk_dummy'],
    'vnnva':       ['deployable_dshk_dummy'],
    'gb82':        ['deployable_stinger_dummy'],
    'arg82':       ['deployable_djigit_dummy'],
    'fr':          ['deployable_mistral_dummy'],
    'nl':          ['deployable_stinger_dummy'],
}
#
#
# HMG SETTINGS
#
# The HMGs templates for each team
# Default is deployable_50cal_tripod_m2, deployable_50cal_tripod_type85, deployable_50cal_tripod_kord, deployable_50cal_tripod_dshk
C['HMG_TEMPLATES'] = {
    'us':          ['deployable_50cal_tripod_m2'],
    'usa':         ['deployable_50cal_tripod_m2'],
    'cf':          ['deployable_50cal_tripod_m2'],
    'gb':          ['deployable_50cal_tripod_m2'],
    'ger':         ['deployable_50cal_tripod_m2'],
    'ch':          ['deployable_50cal_tripod_type85'],
    'mec':         ['deployable_50cal_tripod_kord'],
    'pl':          ['deployable_50cal_tripod_kord'],
    'ru':          ['deployable_50cal_tripod_kord'],
    'idf':         ['deployable_50cal_tripod_m2'],
    'fsa':         ['deployable_50cal_tripod_dshk'],
    'chinsurgent': ['deployable_50cal_tripod_dshk'],
    'vnusa':       ['deployable_50cal_tripod_m2'],
    'vnusmc':      ['deployable_50cal_tripod_m2'],
    'vnnva':       ['deployable_50cal_tripod_dshk'],
    'gb82':        ['deployable_50cal_tripod_m2'],
    'arg82':       ['deployable_50cal_tripod_m2'],
    'fr':          ['deployable_50cal_tripod_m2'],
    'nl':          ['deployable_50cal_tripod_m2'],
    'ww2ger':      ['deployable_mg42'],
    'ww2usa':      ['deployable_m1919a6'],
}
#
#
# The HMGs dummy templates for each team
# Default is deployable_50cal_tripod_m2_dummy, deployable_50cal_tripod_type85_dummy, deployable_50cal_tripod_kord_dummy, deployable_50cal_tripod_dshk_dummy
C['HMG_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_50cal_tripod_m2_dummy'],
    'usa':         ['deployable_50cal_tripod_m2_dummy'],
    'cf':          ['deployable_50cal_tripod_m2_dummy'],
    'gb':          ['deployable_50cal_tripod_m2_dummy'],
    'ger':         ['deployable_50cal_tripod_m2_dummy'],
    'ch':          ['deployable_50cal_tripod_type85_dummy'],
    'mec':         ['deployable_50cal_tripod_kord_dummy'],
    'pl':          ['deployable_50cal_tripod_kord_dummy'],
    'ru':          ['deployable_50cal_tripod_kord_dummy'],
    'idf':         ['deployable_50cal_tripod_m2_dummy'],
    'fsa':         ['deployable_50cal_tripod_dshk_dummy'],
    'chinsurgent': ['deployable_50cal_tripod_dshk_dummy'],
    'vnusa':       ['deployable_50cal_tripod_m2_dummy'],
    'vnusmc':      ['deployable_50cal_tripod_m2_dummy'],
    'vnnva':       ['deployable_50cal_tripod_dshk_dummy'],
    'gb82':        ['deployable_50cal_tripod_m2_dummy'],
    'arg82':       ['deployable_50cal_tripod_m2_dummy'],
    'fr':          ['deployable_50cal_tripod_m2_dummy'],
    'nl':          ['deployable_50cal_tripod_m2_dummy'],
    'ww2ger':      ['deployable_mg42_dummy'],
    'ww2usa':      ['deployable_m1919a6_dummy'],
}
#
#
# TOW SETTINGS
#
# The TOWs templates for each team
# Default is deployable_tow, deployable_hj8, deployable_milan, deployable_spg9
C['TOW_TEMPLATES'] = {
    'us':          ['deployable_tow'],
    'usa':         ['deployable_tow'],
    'cf':          ['deployable_tow'],
    'gb':          ['deployable_tow'],
    'ger':         ['deployable_milan_mira'],
    'ch':          ['deployable_hj8'],
    'mec':         ['deployable_milan'],
    'pl':          ['deployable_spike'],
    'ru':          ['deployable_kornet'],
    'idf':         ['deployable_tow'],
    'arf':         ['deployable_spg9'],
    'chinsurgent': ['deployable_spg9'],
    'fsa':         ['deployable_hj8', 'deployable_milan', 'deployable_tow', 'deployable_kornet'],
    'taliban':     ['deployable_spg9'],
    'hamas':       ['deployable_spg9'],
    'meinsurgent': ['deployable_spg9'],
    'vnnva':       ['deployable_spg9'],
    'gb82':        ['deployable_milan'],
    'arg82':       ['deployable_spg9'],
    'fr':          ['deployable_milan_mira'],
    'nl':          ['deployable_spike'],
}
#
#
# The TOWs dummy templates for each team
# Default is deployable_tow_dummy, deployable_hj8_dummy, deployable_milan_dummy, deployable_spg9_dummy
C['TOW_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_tow_dummy'],
    'usa':         ['deployable_tow_dummy'],
    'cf':          ['deployable_tow_dummy'],
    'gb':          ['deployable_tow_dummy'],
    'ger':         ['deployable_milan_mira_dummy'],
    'ch':          ['deployable_hj8_dummy'],
    'mec':         ['deployable_milan_dummy'],
    'pl':          ['deployable_spike_dummy'],
    'ru':          ['deployable_kornet_dummy'],
    'idf':         ['deployable_tow_dummy'],
    'arf':         ['deployable_spg9_dummy'],
    'fsa':         ['deployable_hj8_dummy'],
    'chinsurgent': ['deployable_spg9_dummy'],
    'taliban':     ['deployable_spg9_dummy'],
    'hamas':       ['deployable_spg9_dummy'],
    'meinsurgent': ['deployable_spg9_dummy'],
    'vnnva':       ['deployable_spg9_dummy'],
    'gb82':        ['deployable_milan_dummy'],
    'arg82':       ['deployable_spg9_dummy'],
    'fr':          ['deployable_milan_mira_dummy'],
    'nl':          ['deployable_spike_dummy'],
}
#
#
# MORTAR SETTINGS
#
# The Mortars templates for each team
# Default is deployable_mortar_m252
C['MORTAR_TEMPLATES'] = {
    'us':          ['deployable_mortar_m252'],
    'usa':         ['deployable_mortar_m252'],
    'cf':          ['deployable_mortar_m252'],
    'gb':          ['deployable_mortar_m252'],
    'ger':         ['deployable_mortar_m252'],
    'ch':          ['deployable_mortar_pp87'],
    'mec':         ['deployable_mortar_m252'],
    'pl':          ['deployable_mortar_m252'],
    'ru':          ['deployable_mortar_2b141_podnos'],
    'idf':         ['deployable_mortar_m252'],
    'arf':         [],
    'chinsurgent': ['deployable_mortar_2b141_podnos'],
    'taliban':     ['deployable_mortar_m252_ins'],
    'hamas':       ['deployable_mortar_m252_ins'],
    'meinsurgent': ['deployable_mortar_m252_ins'],
    'fsa':         ['deployable_mortar_m252_ins'],
    'vnusa':       ['deployable_mortar_m252'],
    'vnusmc':      ['deployable_mortar_m252'],
    'vnnva':       ['deployable_mortar_m252'],
    'gb82':        ['deployable_mortar_m252'],
    'arg82':       ['deployable_mortar_m252'],
    'fr':          ['deployable_mortar_m252'],
    'nl':          ['deployable_mortar_m252'],
    'ww2ger':      [],
    'ww2usa':      [],
}
#
#
# The TOWs dummy templates for each team
# Default is deployable_mortar_m252_dummy
C['MORTAR_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_mortar_m252_dummy'],
    'usa':         ['deployable_mortar_m252_dummy'],
    'cf':          ['deployable_mortar_m252_dummy'],
    'gb':          ['deployable_mortar_m252_dummy'],
    'ger':         ['deployable_mortar_m252_dummy'],
    'ch':          ['deployable_mortar_m252_dummy'],
    'mec':         ['deployable_mortar_m252_dummy'],
    'pl':          ['deployable_mortar_m252_dummy'],
    'ru':          ['deployable_mortar_m252_dummy'],
    'idf':         ['deployable_mortar_m252_dummy'],
    'arf':         ['deployable_mortar_m252_ins_dummy'],
    'chinsurgent': ['deployable_mortar_m252_dummy'],
    'taliban':     ['deployable_mortar_m252_ins_dummy'],
    'hamas':       ['deployable_mortar_m252_ins_dummy'],
    'meinsurgent': ['deployable_mortar_m252_ins_dummy'],
    'fsa':         ['deployable_mortar_m252_ins_dummy'],
    'vnusa':       ['deployable_mortar_m252_dummy'],
    'vnusmc':      ['deployable_mortar_m252_dummy'],
    'vnnva':       ['deployable_mortar_m252_dummy'],
    'gb82':        ['deployable_mortar_m252_dummy'],
    'arg82':       ['deployable_mortar_m252_dummy'],
    'fr':          ['deployable_mortar_m252_dummy'],
    'nl':          ['deployable_mortar_m252_dummy'],
    'ww2ger':      [],
    'ww2usa':      [],
}
#
#
# FOXHOLE SETTINGS
#
# The foxhole templates for each team
# Default is deployable_foxhole
C['FOXHOLE_TEMPLATES'] = {
    'us':          ['deployable_foxhole'],
    'usa':         ['deployable_foxhole'],
    'cf':          ['deployable_foxhole'],
    'gb':          ['deployable_foxhole'],
    'ger':         ['deployable_foxhole'],
    'ch':          ['deployable_foxhole'],
    'mec':         ['deployable_foxhole'],
    'pl':          ['deployable_foxhole'],
    'ru':          ['deployable_foxhole'],
    'idf':         ['deployable_foxhole'],
    'chinsurgent': ['deployable_foxhole'],
    'fsa':         ['deployable_foxhole'],
    'vnusa':       ['deployable_foxhole'],
    'vnusmc':      ['deployable_foxhole'],
    'vnnva':       ['deployable_foxhole'],
    'gb82':        ['deployable_foxhole'],
    'arg82':       ['deployable_foxhole'],
    'fr':          ['deployable_foxhole'],
    'nl':          ['deployable_foxhole'],
    'ww2ger':      ['deployable_foxhole'],
    'ww2usa':      ['deployable_foxhole'],
}
#
#
# The foxhole dummy templates for each team
# Default is deployable_foxhole_dummy
C['FOXHOLE_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_foxhole_dummy'],
    'usa':         ['deployable_foxhole_dummy'],
    'cf':          ['deployable_foxhole_dummy'],
    'gb':          ['deployable_foxhole_dummy'],
    'ger':         ['deployable_foxhole_dummy'],
    'ch':          ['deployable_foxhole_dummy'],
    'mec':         ['deployable_foxhole_dummy'],
    'pl':          ['deployable_foxhole_dummy'],
    'ru':          ['deployable_foxhole_dummy'],
    'idf':         ['deployable_foxhole_dummy'],
    'fsa':         ['deployable_foxhole_dummy'],
    'chinsurgent': ['deployable_foxhole_dummy'],
    'vnusa':       ['deployable_foxhole_dummy'],
    'vnusmc':      ['deployable_foxhole_dummy'],
    'vnnva':       ['deployable_foxhole_dummy'],
    'gb82':        ['deployable_foxhole_dummy'],
    'arg82':       ['deployable_foxhole_dummy'],
    'fr':          ['deployable_foxhole_dummy'],
    'nl':          ['deployable_foxhole_dummy'],
    'ww2ger':      ['deployable_foxhole_dummy'],
    'ww2usa':      ['deployable_foxhole_dummy'],
}
#
#
# RAZORWIRE SETTINGS
#
# The razorwire templates for each team
# Default is deployable_razorwire
C['RAZORWIRES_TEMPLATES'] = {
    'us':          ['deployable_razorwire'],
    'usa':         ['deployable_razorwire'],
    'cf':          ['deployable_razorwire'],
    'gb':          ['deployable_razorwire'],
    'ger':         ['deployable_razorwire'],
    'ch':          ['deployable_razorwire'],
    'mec':         ['deployable_razorwire'],
    'pl':          ['deployable_razorwire'],
    'ru':          ['deployable_razorwire'],
    'idf':         ['deployable_razorwire'],
    'chinsurgent': ['deployable_razorwire'],
    'vnusa':       ['deployable_razorwire'],
    'vnusmc':      ['deployable_razorwire'],
    'vnnva':       ['deployable_razorwire'],
    'gb82':        ['deployable_razorwire'],
    'arg82':       ['deployable_razorwire'],
    'fr':          ['deployable_razorwire'],
    'nl':          ['deployable_razorwire'],
    'ww2ger':      ['deployable_razorwire'],
    'ww2usa':      ['deployable_razorwire'],
}
#
#
# The razorwire templates for each team
# Default is deployable_razorwire_dummy
C['RAZORWIRES_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_razorwire_dummy'],
    'usa':         ['deployable_razorwire_dummy'],
    'cf':          ['deployable_razorwire_dummy'],
    'gb':          ['deployable_razorwire_dummy'],
    'ger':         ['deployable_razorwire_dummy'],
    'ch':          ['deployable_razorwire_dummy'],
    'mec':         ['deployable_razorwire_dummy'],
    'pl':          ['deployable_razorwire_dummy'],
    'ru':          ['deployable_razorwire_dummy'],
    'idf':         ['deployable_razorwire_dummy'],
    'chinsurgent': ['deployable_razorwire_dummy'],
    'vnusa':       ['deployable_razorwire_dummy'],
    'vnusmc':      ['deployable_razorwire_dummy'],
    'vnnva':       ['deployable_razorwire_dummy'],
    'gb82':        ['deployable_razorwire_dummy'],
    'arg82':       ['deployable_razorwire_dummy'],
    'fr':          ['deployable_razorwire_dummy'],
    'nl':          ['deployable_razorwire_dummy'],
    'ww2ger':      ['deployable_razorwire_dummy'],
    'ww2usa':      ['deployable_razorwire_dummy'],
}
#
#
# ROADBLOCK SHORT SETTINGS
#
# The roadblock_short templates for each team
# Default is deployable_razorwire
C['ROADBLOCK_SHORT_TEMPLATES'] = {
    'fsa':         ['deployable_roadblock_s01', 'deployable_roadblock_s02'],
    'hamas':       ['deployable_roadblock_s01', 'deployable_roadblock_s02'],
    'meinsurgent': ['deployable_roadblock_s01', 'deployable_roadblock_s02'],
    'arf':         ['deployable_roadblock_s01', 'deployable_roadblock_s02'],
}
#
#
# The roadblock_short templates for each team
# Default is deployable_roadblock_s_dummy
C['ROADBLOCK_SHORT_TEMPLATES_DUMMY'] = {
    'fsa':         ['deployable_roadblock_short_dummy'],
    'hamas':       ['deployable_roadblock_short_dummy'],
    'meinsurgent': ['deployable_roadblock_short_dummy'],
    'arf':         ['deployable_roadblock_short_dummy'],
}
#
#
# ROADBLOCK LONG SETTINGS
#
# The roadblock_long templates for each team
# Default are 'deployable_roadblock_l01', 'deployable_roadblock_l02', 'deployable_roadblock_l03'
C['ROADBLOCK_LONG_TEMPLATES'] = {
    'fsa':         ['deployable_roadblock_l01', 'deployable_roadblock_l02', 'deployable_roadblock_l03'],
    'hamas':       ['deployable_roadblock_l01', 'deployable_roadblock_l02', 'deployable_roadblock_l03'],
    'meinsurgent': ['deployable_roadblock_l01', 'deployable_roadblock_l02', 'deployable_roadblock_l03'],
    'arf':         ['deployable_roadblock_l01', 'deployable_roadblock_l02', 'deployable_roadblock_l03'],
}
#
#
# The roadblock_long templates for each team
# Default is deployable_roadblock_s_dummy
C['ROADBLOCK_LONG_TEMPLATES_DUMMY'] = {
    'fsa':         ['deployable_roadblock_long_dummy'],
    'hamas':       ['deployable_roadblock_long_dummy'],
    'meinsurgent': ['deployable_roadblock_long_dummy'],
    'arf':         ['deployable_roadblock_long_dummy'],
}
#
#
# SANDBAG WALLS SETTINGS
#
# The sandbags templates for each team
# Default is deployable_sandbags_5m
C['SANDBAGS_TEMPLATES'] = {
    'us':          ['deployable_sandbags_5m'],
    'usa':         ['deployable_sandbags_5m'],
    'cf':          ['deployable_sandbags_5m'],
    'gb':          ['deployable_sandbags_5m'],
    'ger':         ['deployable_sandbags_5m'],
    'ch':          ['deployable_sandbags_5m'],
    'mec':         ['deployable_sandbags_5m'],
    'pl':          ['deployable_sandbags_5m'],
    'ru':          ['deployable_sandbags_5m'],
    'idf':         ['deployable_sandbags_5m'],
    'fsa':         ['deployable_sandbags_5m'],
    'chinsurgent': ['deployable_sandbags_5m'],
    'vnusa':       ['deployable_sandbags_5m'],
    'vnusmc':      ['deployable_sandbags_5m'],
    'vnnva':       ['deployable_sandbags_5m'],
    'gb82':        ['deployable_sandbags_5m'],
    'arg82':       ['deployable_sandbags_5m'],
    'fr':          ['deployable_sandbags_5m'],
    'nl':          ['deployable_sandbags_5m'],
    'ww2ger':      ['deployable_sandbags_5m'],
    'ww2usa':      ['deployable_sandbags_5m'],
}
#
#
# The sandbags dummy templates for each team
# Default is deployable_sandbags_5m_dummy
C['SANDBAGS_TEMPLATES_DUMMY'] = {
    'us':          ['deployable_sandbags_5m_dummy'],
    'usa':         ['deployable_sandbags_5m_dummy'],
    'cf':          ['deployable_sandbags_5m_dummy'],
    'gb':          ['deployable_sandbags_5m_dummy'],
    'ger':         ['deployable_sandbags_5m_dummy'],
    'ch':          ['deployable_sandbags_5m_dummy'],
    'mec':         ['deployable_sandbags_5m_dummy'],
    'pl':          ['deployable_sandbags_5m_dummy'],
    'ru':          ['deployable_sandbags_5m_dummy'],
    'idf':         ['deployable_sandbags_5m_dummy'],
    'fsa':         ['deployable_sandbags_5m_dummy'],
    'chinsurgent': ['deployable_sandbags_5m_dummy'],
    'vnusa':       ['deployable_sandbags_5m_dummy'],
    'vnusmc':      ['deployable_sandbags_5m_dummy'],
    'vnnva':       ['deployable_sandbags_5m_dummy'],
    'gb82':        ['deployable_sandbags_5m_dummy'],
    'arg82':       ['deployable_sandbags_5m_dummy'],
    'fr':          ['deployable_sandbags_5m_dummy'],
    'nl':          ['deployable_sandbags_5m_dummy'],
    'ww2ger':      ['deployable_sandbags_5m_dummy'],
    'ww2usa':      ['deployable_sandbags_5m_dummy'],
}
#
#
# PROJECT REALITY CIVILIAN CLASS SETTINGS
#
# Number of seconds added to the player for each civilian kill
# Default is 120 seconds
C['CIV_PENALTY_PER_COUNT'] = 120
#
#
# Max number of seconds for the penalty
# Default is 3 times the CIV_PENALTY_PER_COUNT
C['CIV_MAX_PENALTY'] = 3 * C['CIV_PENALTY_PER_COUNT']
#
#
# Number of seconds added to the civilian when he gets captured
# Default is 90 seconds
C['CIV_CAPTURE_PENALTY'] = 90
#
# Number of seconds the civilian stays inside the ROE after helping insurgents
# Default is 60 seconds
C['CIV_HELP_INTERVAL'] = 120
#
# Distance the civilian has to stay away from armed insurgents to not be inside the ROE for helping insurgents (horizontal)
# Default is 5 meters
C['CIV_HELP_DISTANCE_XZ'] = 5
#
# Distance the civilian has to stay away from armed insurgents to not be inside the ROE for helping insurgents (vertical)
# Default is 1.5 meters
C['CIV_HELP_DISTANCE_Y'] = 1.5
#
#
# Teams that have civilian logic
# Default is meinsurgent, hamas
C['CIV_TEAMS'] = ['meinsurgent', 'hamas', 'taliban']
#
#
# Civilian kit name
# Default is unarmed, meinsurgent_medic_alt
C['CIV_KIT_NAME'] = ['hamas_unarmed', 'meinsurgent_unarmed', 'meinsurgent_medic_alt', 'taliban_unarmed']
#
#
# Weapon prefix that arrests civilians
# Default is ziptie
C['CIV_CAPTURE_WEAPON_TEMPLATES'] = ['ziptie']
#
#
# Weapon prefix that arrests civilians
# Default is WEAPON_TYPE_SHOTGUN
C['CIV_CAPTURE_WEAPON_TYPES'] = [CONSTANTS.WEAPON_TYPE_SHOTGUN]
#
#
# Number of points added to the attacker that captured a civilian
# Default is 100 points
C['CIV_CAPTURE_SCORE'] = 100
#
#
# Weapon prefix that permits attackers to kill the civilian inside the ROE
# Default is klappspaten
C['CIV_KILL_WEAPON_TEMPLATES'] = ['klappspaten']
#
#
# Number of points added to the attacker that killed a civilian
# Default is -100 points
C['CIV_KILL_SCORE'] = -100
#
#
# Number of kills removed from the attacker count
# Default is -1 (the civilian doesn't count as a kill)
C['CIV_KILL_COUNT_PENALTY'] = -1
#
#
# Number of seconds the attacker stays with blocked kit request
# Default is 480 seconds (8 minutes)
C['CIV_KILL_PENALTY'] = 480
#
#
# Maximum number of civilians kills before being arrested
# Default is 5
C['CIV_KILL_MAX'] = 5
#
#
# Interval where civilian kills are removed from the player history
# Default is 900 seconds
C['CIV_KILL_INTERVAL'] = 900
#
#
# PROJECT REALITY INFORMANTS SETTINGS
#
# Factions that have informants
# Default are: meinsurgent, taliban, hamas, arf, fsa
C['INFORMANTS_TEAMS'] = ['meinsurgent', 'taliban', 'hamas', 'arf', 'fsa']
#
#
# Number of enemies that need to be close so the commander is informed.
# Default is 4
C['INFORMANTS_CLOSE_DISABLE'] = 4
#
#
# How many seconds it takes for the informant to get to the marked position.
# Default is 60 seconds
C['INFORMANTS_MOVEMENT_DELAY'] = 60
#
#
# How many seconds it takes for the intel to reach the commander.
# Default is 15 seconds
C['INFORMANTS_REPORT_DELAY'] = 15
#
#
# PROJECT REALITY MUTINY SETTINGS
#
# Interval between mutinies
# Default is 300 seconds
C['MUTINY_INTERVAL'] = 300
#
#
# Interval for voting
# Default is 30 seconds
C['MUTINY_VOTING'] = 30
#
#
# Percentage for successfully voting a commander off
# Default is 0.5
C['MUTINY_PERCENTAGE'] = 0.5
#
#
# PROJECT REALITY MISC SETTINGS
#
# Team names
C['TEAM_NAME'] = {
    "us":          "The USMC Forces",
    "usa":         "The US Army Forces",
    "cf":          "The Canadian Forces",
    "ch":          "The Chinese Forces",
    "mec":         "The MEC Forces",
    "gb":          "The British Forces",
    "ger":         "The German Forces",
    "pl":          "The Polish Forces",
    "ru":          "The Russian Forces",
    "idf":         "The Israel Defense Forces",
    "arf":         "The African Resistance Fighters",
    "meinsurgent": "The Insurgency",
    "chinsurgent": "The Militia",
    "taliban":     "The Taliban",
    "fsa":         "The Syrian Rebels",
    "hamas":       "The Hamas",
    "vnusa":       "The US Army Forces",
    "vnusmc":      "The USMC Forces",
    "vnnva":       "The North Vietnamese Army",
    "gb82":        "The British Forces",
    "arg82":       "The Argentinian Forces",
    "fr":          "The French Forces",
    "nl":          "The Dutch Armed Forces",
    "ww2ger":      "Wehrmacht",
    "ww2usa":      "The US Army Forces"
}

# List of projectile templates that ignore explosion lag compensation.
# Put server-guided and server-triggered projectiles here (Mines, AA)
C['EXPLOSION_LAGCOMP_DISABLE'] = [
    # Ground to Air
    "stinger_missile", "sa7_missile", "sa3_missile", "sa19_grison", "ru_aa_9M333", "igla_9k38", "gbaa_blowpipe", "seacat_missile",
    # Air to air
    "aa11_archer", "aim9", "aim120_amraam", "ru_aa_r60m", "aim9b_sidewinder",
    "matra_r550_magic1", "matra_r530", "aim9l_sidewinder", "ru_aa_r27re",
    # SACLOS -
    # Note that StarStreak anti tank damage is based on direct hits, so its okay to not compensate the explosion
    "sa19_grison_saclos", "starstreak_hvm", "seacat_missile_aclos",
    # Mines
    "at_mine_projectile", "at_mine_tm35_projectile", "rumin_mon50_Projectile",
    "usmin_m1a1_projectile", "usmin_m2a3_projectile", "vnhgr_betty_Projectile",
]
#
#
# Defines if all players in a server running the private config have debug powers
# Default is 0 (no)
C['PRDEBUG_ALL'] = 0
#
#
# ========================================================================================================
