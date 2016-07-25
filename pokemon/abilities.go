// Copyright (c) 2011-2016 Guangcong Luo and other contributors
// http://pokemonshowdown.com/
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package pokemon

import "github.com/xfix/showdown2irc/showdown"

var abilities = map[showdown.UserID]*Ability{
	"adaptability": {
		Name:        "Adaptability",
		Description: "This Pokemon's same-type attack bonus (STAB) is 2 instead of 1.5.",
	},
	"aftermath": {
		Name:        "Aftermath",
		Description: "If this Pokemon is KOed with a contact move, that move's user loses 1/4 its max HP.",
	},
	"aerilate": {
		Name:        "Aerilate",
		Description: "This Pokemon's Normal-type moves become Flying type and have 1.3x power.",
	},
	"airlock": {
		Name:        "Air Lock",
		Description: "While this Pokemon is active, the effects of weather conditions are disabled.",
	},
	"analytic": {
		Name:        "Analytic",
		Description: "This Pokemon's attacks have 1.3x power if it is the last to move in a turn.",
	},
	"angerpoint": {
		Name:        "Anger Point",
		Description: "If this Pokemon (not its substitute) takes a critical hit, its Attack is raised 12 stages.",
	},
	"anticipation": {
		Name:        "Anticipation",
		Description: "On switch-in, this Pokemon shudders if any foe has a supereffective or OHKO move.",
	},
	"arenatrap": {
		Name:        "Arena Trap",
		Description: "Prevents adjacent foes from choosing to switch unless they are airborne.",
	},
	"aromaveil": {
		Name:        "Aroma Veil",
		Description: "Protects user/allies from Attract, Disable, Encore, Heal Block, Taunt, and Torment.",
	},
	"aurabreak": {
		Name:        "Aura Break",
		Description: "While this Pokemon is active, the Dark Aura and Fairy Aura power modifier is 0.75x.",
	},
	"baddreams": {
		Name:        "Bad Dreams",
		Description: "Causes sleeping adjacent foes to lose 1/8 of their max HP at the end of each turn.",
	},
	"battlearmor": {
		Name:        "Battle Armor",
		Description: "This Pokemon cannot be struck by a critical hit.",
	},
	"bigpecks": {
		Name:        "Big Pecks",
		Description: "Prevents other Pokemon from lowering this Pokemon's Defense stat stage.",
	},
	"blaze": {
		Name:        "Blaze",
		Description: "When this Pokemon has 1/3 or less of its max HP, its Fire attacks do 1.5x damage.",
	},
	"bulletproof": {
		Name:        "Bulletproof",
		Description: "Makes user immune to ballistic moves (Shadow Ball, Sludge Bomb, Focus Blast, etc).",
	},
	"cheekpouch": {
		Name:        "Cheek Pouch",
		Description: "If this Pokemon eats a Berry, it restores 1/3 of its max HP after the Berry's effect.",
	},
	"chlorophyll": {
		Name:        "Chlorophyll",
		Description: "If Sunny Day is active, this Pokemon's Speed is doubled.",
	},
	"clearbody": {
		Name:        "Clear Body",
		Description: "Prevents other Pokemon from lowering this Pokemon's stat stages.",
	},
	"cloudnine": {
		Name:        "Cloud Nine",
		Description: "While this Pokemon is active, the effects of weather conditions are disabled.",
	},
	"colorchange": {
		Name:        "Color Change",
		Description: "This Pokemon's type changes to the type of a move it's hit by, unless it has the type.",
	},
	"competitive": {
		Name:        "Competitive",
		Description: "This Pokemon's Sp. Atk is raised by 2 for each of its stats that is lowered by a foe.",
	},
	"compoundeyes": {
		Name:        "Compound Eyes",
		Description: "This Pokemon's moves have their accuracy multiplied by 1.3.",
	},
	"contrary": {
		Name:        "Contrary",
		Description: "If this Pokemon has a stat stage raised it is lowered instead, and vice versa.",
	},
	"cursedbody": {
		Name:        "Cursed Body",
		Description: "If this Pokemon is hit by an attack, there is a 30% chance that move gets disabled.",
	},
	"cutecharm": {
		Name:        "Cute Charm",
		Description: "30% chance of infatuating Pokemon of the opposite gender if they make contact.",
	},
	"damp": {
		Name:        "Damp",
		Description: "While this Pokemon is active, Self-Destruct, Explosion, and Aftermath have no effect.",
	},
	"darkaura": {
		Name:        "Dark Aura",
		Description: "While this Pokemon is active, a Dark move used by any Pokemon has 1.33x power.",
	},
	"defeatist": {
		Name:        "Defeatist",
		Description: "While this Pokemon has 1/2 or less of its max HP, its Attack and Sp. Atk are halved.",
	},
	"defiant": {
		Name:        "Defiant",
		Description: "This Pokemon's Attack is raised by 2 for each of its stats that is lowered by a foe.",
	},
	"deltastream": {
		Name:        "Delta Stream",
		Description: "On switch-in, strong winds begin until this Ability is not active in battle.",
	},
	"desolateland": {
		Name:        "Desolate Land",
		Description: "On switch-in, extremely harsh sunlight begins until this Ability is not active in battle.",
	},
	"download": {
		Name:        "Download",
		Description: "On switch-in, Attack or Sp. Atk is raised 1 stage based on the foes' weaker Defense.",
	},
	"drizzle": {
		Name:        "Drizzle",
		Description: "On switch-in, this Pokemon summons Rain Dance.",
	},
	"drought": {
		Name:        "Drought",
		Description: "On switch-in, this Pokemon summons Sunny Day.",
	},
	"dryskin": {
		Name:        "Dry Skin",
		Description: "This Pokemon is healed 1/4 by Water, 1/8 by Rain; is hurt 1.25x by Fire, 1/8 by Sun.",
	},
	"earlybird": {
		Name:        "Early Bird",
		Description: "This Pokemon's sleep counter drops by 2 instead of 1.",
	},
	"effectspore": {
		Name:        "Effect Spore",
		Description: "30% chance of poison/paralysis/sleep on others making contact with this Pokemon.",
	},
	"fairyaura": {
		Name:        "Fairy Aura",
		Description: "While this Pokemon is active, a Fairy move used by any Pokemon has 1.33x power.",
	},
	"filter": {
		Name:        "Filter",
		Description: "This Pokemon receives 3/4 damage from supereffective attacks.",
	},
	"flamebody": {
		Name:        "Flame Body",
		Description: "30% chance a Pokemon making contact with this Pokemon will be burned.",
	},
	"flareboost": {
		Name:        "Flare Boost",
		Description: "While this Pokemon is burned, its special attacks have 1.5x power.",
	},
	"flashfire": {
		Name:        "Flash Fire",
		Description: "This Pokemon's Fire attacks do 1.5x damage if hit by one Fire move; Fire immunity.",
	},
	"flowergift": {
		Name:        "Flower Gift",
		Description: "If user is Cherrim and Sunny Day is active, it and allies' Attack and Sp. Def are 1.5x.",
	},
	"flowerveil": {
		Name:        "Flower Veil",
		Description: "This side's Grass types can't have stats lowered or status inflicted by other Pokemon.",
	},
	"forecast": {
		Name:        "Forecast",
		Description: "Castform's type changes to the current weather condition's type, except Sandstorm.",
	},
	"forewarn": {
		Name:        "Forewarn",
		Description: "On switch-in, this Pokemon is alerted to the foes' move with the highest power.",
	},
	"friendguard": {
		Name:        "Friend Guard",
		Description: "This Pokemon's allies receive 3/4 damage from other Pokemon's attacks.",
	},
	"frisk": {
		Name:        "Frisk",
		Description: "On switch-in, this Pokemon identifies the held items of all opposing Pokemon.",
	},
	"furcoat": {
		Name:        "Fur Coat",
		Description: "This Pokemon's Defense is doubled.",
	},
	"galewings": {
		Name:        "Gale Wings",
		Description: "This Pokemon's Flying-type moves have their priority increased by 1.",
	},
	"gluttony": {
		Name:        "Gluttony",
		Description: "When this Pokemon has 1/2 or less of its maximum HP, it uses certain Berries early.",
	},
	"gooey": {
		Name:        "Gooey",
		Description: "Pokemon making contact with this Pokemon have their Speed lowered by 1 stage.",
	},
	"grasspelt": {
		Name:        "Grass Pelt",
		Description: "If Grassy Terrain is active, this Pokemon's Defense is multiplied by 1.5.",
	},
	"guts": {
		Name:        "Guts",
		Description: "If this Pokemon is statused, its Attack is 1.5x; ignores burn halving physical damage.",
	},
	"harvest": {
		Name:        "Harvest",
		Description: "If last item used is a Berry, 50% chance to restore it each end of turn. 100% in Sun.",
	},
	"healer": {
		Name:        "Healer",
		Description: "30% chance of curing an adjacent ally's status at the end of each turn.",
	},
	"heatproof": {
		Name:        "Heatproof",
		Description: "The power of Fire-type attacks against this Pokemon is halved; burn damage halved.",
	},
	"heavymetal": {
		Name:        "Heavy Metal",
		Description: "This Pokemon's weight is doubled.",
	},
	"honeygather": {
		Name:        "Honey Gather",
		Description: "No competitive use.",
	},
	"hugepower": {
		Name:        "Huge Power",
		Description: "This Pokemon's Attack is doubled.",
	},
	"hustle": {
		Name:        "Hustle",
		Description: "This Pokemon's Attack is 1.5x and accuracy of its physical attacks is 0.8x.",
	},
	"hydration": {
		Name:        "Hydration",
		Description: "This Pokemon has its status cured at the end of each turn if Rain Dance is active.",
	},
	"hypercutter": {
		Name:        "Hyper Cutter",
		Description: "Prevents other Pokemon from lowering this Pokemon's Attack stat stage.",
	},
	"icebody": {
		Name:        "Ice Body",
		Description: "If Hail is active, this Pokemon heals 1/16 of its max HP each turn; immunity to Hail.",
	},
	"illuminate": {
		Name:        "Illuminate",
		Description: "No competitive use.",
	},
	"illusion": {
		Name:        "Illusion",
		Description: "This Pokemon appears as the last Pokemon in the party until it takes direct damage.",
	},
	"immunity": {
		Name:        "Immunity",
		Description: "This Pokemon cannot be poisoned. Gaining this Ability while poisoned cures it.",
	},
	"imposter": {
		Name:        "Imposter",
		Description: "On switch-in, this Pokemon Transforms into the opposing Pokemon that is facing it.",
	},
	"infiltrator": {
		Name:        "Infiltrator",
		Description: "Moves ignore substitutes and opposing Reflect, Light Screen, Safeguard, and Mist.",
	},
	"innerfocus": {
		Name:        "Inner Focus",
		Description: "This Pokemon cannot be made to flinch.",
	},
	"insomnia": {
		Name:        "Insomnia",
		Description: "This Pokemon cannot fall asleep. Gaining this Ability while asleep cures it.",
	},
	"intimidate": {
		Name:        "Intimidate",
		Description: "On switch-in, this Pokemon lowers the Attack of adjacent opponents by 1 stage.",
	},
	"ironbarbs": {
		Name:        "Iron Barbs",
		Description: "Pokemon making contact with this Pokemon lose 1/8 of their max HP.",
	},
	"ironfist": {
		Name:        "Iron Fist",
		Description: "This Pokemon's punch-based attacks have 1.2x power. Sucker Punch is not boosted.",
	},
	"justified": {
		Name:        "Justified",
		Description: "This Pokemon's Attack is raised by 1 stage after it is damaged by a Dark-type move.",
	},
	"keeneye": {
		Name:        "Keen Eye",
		Description: "This Pokemon's accuracy can't be lowered by others; ignores their evasiveness stat.",
	},
	"klutz": {
		Name:        "Klutz",
		Description: "This Pokemon's held item has no effect, except Macho Brace. Fling cannot be used.",
	},
	"leafguard": {
		Name:        "Leaf Guard",
		Description: "If Sunny Day is active, this Pokemon cannot be statused and Rest will fail for it.",
	},
	"levitate": {
		Name:        "Levitate",
		Description: "This Pokemon is immune to Ground; Gravity/Ingrain/Smack Down/Iron Ball nullify it.",
	},
	"lightmetal": {
		Name:        "Light Metal",
		Description: "This Pokemon's weight is halved.",
	},
	"lightningrod": {
		Name:        "Lightning Rod",
		Description: "This Pokemon draws Electric moves to itself to raise Sp. Atk by 1; Electric immunity.",
	},
	"limber": {
		Name:        "Limber",
		Description: "This Pokemon cannot be paralyzed. Gaining this Ability while paralyzed cures it.",
	},
	"liquidooze": {
		Name:        "Liquid Ooze",
		Description: "This Pokemon damages those draining HP from it for as much as they would heal.",
	},
	"magicbounce": {
		Name:        "Magic Bounce",
		Description: "This Pokemon blocks certain status moves and bounces them back to the user.",
	},
	"magicguard": {
		Name:        "Magic Guard",
		Description: "This Pokemon can only be damaged by direct attacks.",
	},
	"magician": {
		Name:        "Magician",
		Description: "If this Pokemon has no item, it steals the item off a Pokemon it hits with an attack.",
	},
	"magmaarmor": {
		Name:        "Magma Armor",
		Description: "This Pokemon cannot be frozen. Gaining this Ability while frozen cures it.",
	},
	"magnetpull": {
		Name:        "Magnet Pull",
		Description: "Prevents adjacent Steel-type foes from choosing to switch.",
	},
	"marvelscale": {
		Name:        "Marvel Scale",
		Description: "If this Pokemon is statused, its Defense is 1.5x.",
	},
	"megalauncher": {
		Name:        "Mega Launcher",
		Description: "This Pokemon's pulse moves have 1.5x power. Heal Pulse heals 3/4 target's max HP.",
	},
	"minus": {
		Name:        "Minus",
		Description: "If an active ally has this Ability or the Ability Plus, this Pokemon's Sp. Atk is 1.5x.",
	},
	"moldbreaker": {
		Name:        "Mold Breaker",
		Description: "This Pokemon's moves and their effects ignore the Abilities of other Pokemon.",
	},
	"moody": {
		Name:        "Moody",
		Description: "Raises a random stat by 2 and lowers another stat by 1 at the end of each turn.",
	},
	"motordrive": {
		Name:        "Motor Drive",
		Description: "This Pokemon's Speed is raised 1 stage if hit by an Electric move; Electric immunity.",
	},
	"moxie": {
		Name:        "Moxie",
		Description: "This Pokemon's Attack is raised by 1 stage if it attacks and KOes another Pokemon.",
	},
	"multiscale": {
		Name:        "Multiscale",
		Description: "If this Pokemon is at full HP, damage taken from attacks is halved.",
	},
	"multitype": {
		Name:        "Multitype",
		Description: "If this Pokemon is an Arceus, its type changes to match its held Plate.",
	},
	"mummy": {
		Name:        "Mummy",
		Description: "Pokemon making contact with this Pokemon have their Ability changed to Mummy.",
	},
	"naturalcure": {
		Name:        "Natural Cure",
		Description: "This Pokemon has its major status condition cured when it switches out.",
	},
	"noguard": {
		Name:        "No Guard",
		Description: "Every move used by or against this Pokemon will always hit.",
	},
	"normalize": {
		Name:        "Normalize",
		Description: "This Pokemon's moves are changed to be Normal type.",
	},
	"oblivious": {
		Name:        "Oblivious",
		Description: "This Pokemon cannot be infatuated or taunted. Gaining this Ability cures it.",
	},
	"overcoat": {
		Name:        "Overcoat",
		Description: "This Pokemon is immune to powder moves and damage from Sandstorm or Hail.",
	},
	"overgrow": {
		Name:        "Overgrow",
		Description: "When this Pokemon has 1/3 or less of its max HP, its Grass attacks do 1.5x damage.",
	},
	"owntempo": {
		Name:        "Own Tempo",
		Description: "This Pokemon cannot be confused. Gaining this Ability while confused cures it.",
	},
	"parentalbond": {
		Name:        "Parental Bond",
		Description: "This Pokemon's damaging moves hit twice. The second hit has its damage halved.",
	},
	"pickup": {
		Name:        "Pickup",
		Description: "If this Pokemon has no item, it finds one used by an adjacent Pokemon this turn.",
	},
	"pickpocket": {
		Name:        "Pickpocket",
		Description: "If this Pokemon has no item, it steals the item off a Pokemon making contact with it.",
	},
	"pixilate": {
		Name:        "Pixilate",
		Description: "This Pokemon's Normal-type moves become Fairy type and have 1.3x power.",
	},
	"plus": {
		Name:        "Plus",
		Description: "If an active ally has this Ability or the Ability Minus, this Pokemon's Sp. Atk is 1.5x.",
	},
	"poisonheal": {
		Name:        "Poison Heal",
		Description: "This Pokemon is healed by 1/8 of its max HP each turn when poisoned; no HP loss.",
	},
	"poisonpoint": {
		Name:        "Poison Point",
		Description: "30% chance a Pokemon making contact with this Pokemon will be poisoned.",
	},
	"poisontouch": {
		Name:        "Poison Touch",
		Description: "This Pokemon's contact moves have a 30% chance of poisoning.",
	},
	"prankster": {
		Name:        "Prankster",
		Description: "This Pokemon's non-damaging moves have their priority increased by 1.",
	},
	"pressure": {
		Name:        "Pressure",
		Description: "If this Pokemon is the target of a foe's move, that move loses one additional PP.",
	},
	"primordialsea": {
		Name:        "Primordial Sea",
		Description: "On switch-in, heavy rain begins until this Ability is not active in battle.",
	},
	"protean": {
		Name:        "Protean",
		Description: "This Pokemon's type changes to match the type of the move it is about to use.",
	},
	"purepower": {
		Name:        "Pure Power",
		Description: "This Pokemon's Attack is doubled.",
	},
	"quickfeet": {
		Name:        "Quick Feet",
		Description: "If this Pokemon is statused, its Speed is 1.5x; ignores Speed drop from paralysis.",
	},
	"raindish": {
		Name:        "Rain Dish",
		Description: "If Rain Dance is active, this Pokemon heals 1/16 of its max HP each turn.",
	},
	"rattled": {
		Name:        "Rattled",
		Description: "This Pokemon's Speed is raised 1 stage if hit by a Bug-, Dark-, or Ghost-type attack.",
	},
	"reckless": {
		Name:        "Reckless",
		Description: "This Pokemon's attacks with recoil or crash damage have 1.2x power; not Struggle.",
	},
	"refrigerate": {
		Name:        "Refrigerate",
		Description: "This Pokemon's Normal-type moves become Ice type and have 1.3x power.",
	},
	"regenerator": {
		Name:        "Regenerator",
		Description: "This Pokemon restores 1/3 of its maximum HP, rounded down, when it switches out.",
	},
	"rivalry": {
		Name:        "Rivalry",
		Description: "This Pokemon's attacks do 1.25x on same gender targets; 0.75x on opposite gender.",
	},
	"rockhead": {
		Name:        "Rock Head",
		Description: "This Pokemon does not take recoil damage besides Struggle/Life Orb/crash damage.",
	},
	"roughskin": {
		Name:        "Rough Skin",
		Description: "Pokemon making contact with this Pokemon lose 1/8 of their max HP.",
	},
	"runaway": {
		Name:        "Run Away",
		Description: "No competitive use.",
	},
	"sandforce": {
		Name:        "Sand Force",
		Description: "This Pokemon's Ground/Rock/Steel attacks do 1.3x in Sandstorm; immunity to it.",
	},
	"sandrush": {
		Name:        "Sand Rush",
		Description: "If Sandstorm is active, this Pokemon's Speed is doubled; immunity to Sandstorm.",
	},
	"sandstream": {
		Name:        "Sand Stream",
		Description: "On switch-in, this Pokemon summons Sandstorm.",
	},
	"sandveil": {
		Name:        "Sand Veil",
		Description: "If Sandstorm is active, this Pokemon's evasiveness is 1.25x; immunity to Sandstorm.",
	},
	"sapsipper": {
		Name:        "Sap Sipper",
		Description: "This Pokemon's Attack is raised 1 stage if hit by a Grass move; Grass immunity.",
	},
	"scrappy": {
		Name:        "Scrappy",
		Description: "This Pokemon can hit Ghost types with Normal- and Fighting-type moves.",
	},
	"serenegrace": {
		Name:        "Serene Grace",
		Description: "This Pokemon's moves have their secondary effect chance doubled.",
	},
	"shadowtag": {
		Name:        "Shadow Tag",
		Description: "Prevents adjacent foes from choosing to switch unless they also have this Ability.",
	},
	"shedskin": {
		Name:        "Shed Skin",
		Description: "This Pokemon has a 33% chance to have its status cured at the end of each turn.",
	},
	"sheerforce": {
		Name:        "Sheer Force",
		Description: "This Pokemon's attacks with secondary effects have 1.3x power; nullifies the effects.",
	},
	"shellarmor": {
		Name:        "Shell Armor",
		Description: "This Pokemon cannot be struck by a critical hit.",
	},
	"shielddust": {
		Name:        "Shield Dust",
		Description: "This Pokemon is not affected by the secondary effect of another Pokemon's attack.",
	},
	"simple": {
		Name:        "Simple",
		Description: "If this Pokemon's stat stages are raised or lowered, the effect is doubled instead.",
	},
	"skilllink": {
		Name:        "Skill Link",
		Description: "This Pokemon's multi-hit attacks always hit the maximum number of times.",
	},
	"slowstart": {
		Name:        "Slow Start",
		Description: "On switch-in, this Pokemon's Attack and Speed are halved for 5 turns.",
	},
	"sniper": {
		Name:        "Sniper",
		Description: "If this Pokemon strikes with a critical hit, the damage is multiplied by 1.5.",
	},
	"snowcloak": {
		Name:        "Snow Cloak",
		Description: "If Hail is active, this Pokemon's evasiveness is 1.25x; immunity to Hail.",
	},
	"snowwarning": {
		Name:        "Snow Warning",
		Description: "On switch-in, this Pokemon summons Hail.",
	},
	"solarpower": {
		Name:        "Solar Power",
		Description: "If Sunny Day is active, this Pokemon's Sp. Atk is 1.5x; loses 1/8 max HP per turn.",
	},
	"solidrock": {
		Name:        "Solid Rock",
		Description: "This Pokemon receives 3/4 damage from supereffective attacks.",
	},
	"soundproof": {
		Name:        "Soundproof",
		Description: "This Pokemon is immune to sound-based moves, including Heal Bell.",
	},
	"speedboost": {
		Name:        "Speed Boost",
		Description: "This Pokemon's Speed is raised 1 stage at the end of each full turn on the field.",
	},
	"stall": {
		Name:        "Stall",
		Description: "This Pokemon moves last among Pokemon using the same or greater priority moves.",
	},
	"stancechange": {
		Name:        "Stance Change",
		Description: "If Aegislash, changes Forme to Blade before attacks and Shield before King's Shield.",
	},
	"static": {
		Name:        "Static",
		Description: "30% chance a Pokemon making contact with this Pokemon will be paralyzed.",
	},
	"steadfast": {
		Name:        "Steadfast",
		Description: "If this Pokemon flinches, its Speed is raised by 1 stage.",
	},
	"stench": {
		Name:        "Stench",
		Description: "This Pokemon's attacks without a chance to flinch have a 10% chance to flinch.",
	},
	"stickyhold": {
		Name:        "Sticky Hold",
		Description: "This Pokemon cannot lose its held item due to another Pokemon's attack.",
	},
	"stormdrain": {
		Name:        "Storm Drain",
		Description: "This Pokemon draws Water moves to itself to raise Sp. Atk by 1; Water immunity.",
	},
	"strongjaw": {
		Name:        "Strong Jaw",
		Description: "This Pokemon's bite-based attacks have 1.5x power. Bug Bite is not boosted.",
	},
	"sturdy": {
		Name:        "Sturdy",
		Description: "If this Pokemon is at full HP, it survives one hit with at least 1 HP. Immune to OHKO.",
	},
	"suctioncups": {
		Name:        "Suction Cups",
		Description: "This Pokemon cannot be forced to switch out by another Pokemon's attack or item.",
	},
	"superluck": {
		Name:        "Super Luck",
		Description: "This Pokemon's critical hit ratio is raised by 1 stage.",
	},
	"swarm": {
		Name:        "Swarm",
		Description: "When this Pokemon has 1/3 or less of its max HP, its Bug attacks do 1.5x damage.",
	},
	"sweetveil": {
		Name:        "Sweet Veil",
		Description: "This Pokemon and its allies cannot fall asleep.",
	},
	"swiftswim": {
		Name:        "Swift Swim",
		Description: "If Rain Dance is active, this Pokemon's Speed is doubled.",
	},
	"symbiosis": {
		Name:        "Symbiosis",
		Description: "If an ally uses its item, this Pokemon gives its item to that ally immediately.",
	},
	"synchronize": {
		Name:        "Synchronize",
		Description: "If another Pokemon burns/poisons/paralyzes this Pokemon, it also gets that status.",
	},
	"tangledfeet": {
		Name:        "Tangled Feet",
		Description: "This Pokemon's evasiveness is doubled as long as it is confused.",
	},
	"technician": {
		Name:        "Technician",
		Description: "This Pokemon's moves of 60 power or less have 1.5x power. Includes Struggle.",
	},
	"telepathy": {
		Name:        "Telepathy",
		Description: "This Pokemon does not take damage from attacks made by its allies.",
	},
	"teravolt": {
		Name:        "Teravolt",
		Description: "This Pokemon's moves and their effects ignore the Abilities of other Pokemon.",
	},
	"thickfat": {
		Name:        "Thick Fat",
		Description: "Fire/Ice-type moves against this Pokemon deal damage with a halved attacking stat.",
	},
	"tintedlens": {
		Name:        "Tinted Lens",
		Description: "This Pokemon's attacks that are not very effective on a target deal double damage.",
	},
	"torrent": {
		Name:        "Torrent",
		Description: "When this Pokemon has 1/3 or less of its max HP, its Water attacks do 1.5x damage.",
	},
	"toxicboost": {
		Name:        "Toxic Boost",
		Description: "While this Pokemon is poisoned, its physical attacks have 1.5x power.",
	},
	"toughclaws": {
		Name:        "Tough Claws",
		Description: "This Pokemon's contact moves have their power multiplied by 1.3.",
	},
	"trace": {
		Name:        "Trace",
		Description: "On switch-in, or when it can, this Pokemon copies a random adjacent foe's Ability.",
	},
	"truant": {
		Name:        "Truant",
		Description: "This Pokemon skips every other turn instead of using a move.",
	},
	"turboblaze": {
		Name:        "Turboblaze",
		Description: "This Pokemon's moves and their effects ignore the Abilities of other Pokemon.",
	},
	"unaware": {
		Name:        "Unaware",
		Description: "This Pokemon ignores other Pokemon's stat stages when taking or doing damage.",
	},
	"unburden": {
		Name:        "Unburden",
		Description: "Speed is doubled on held item loss; boost is lost if it switches, gets new item/Ability.",
	},
	"unnerve": {
		Name:        "Unnerve",
		Description: "While this Pokemon is active, it prevents opposing Pokemon from using their Berries.",
	},
	"victorystar": {
		Name:        "Victory Star",
		Description: "This Pokemon and its allies' moves have their accuracy multiplied by 1.1.",
	},
	"vitalspirit": {
		Name:        "Vital Spirit",
		Description: "This Pokemon cannot fall asleep. Gaining this Ability while asleep cures it.",
	},
	"voltabsorb": {
		Name:        "Volt Absorb",
		Description: "This Pokemon heals 1/4 of its max HP when hit by Electric moves; Electric immunity.",
	},
	"waterabsorb": {
		Name:        "Water Absorb",
		Description: "This Pokemon heals 1/4 of its max HP when hit by Water moves; Water immunity.",
	},
	"waterveil": {
		Name:        "Water Veil",
		Description: "This Pokemon cannot be burned. Gaining this Ability while burned cures it.",
	},
	"weakarmor": {
		Name:        "Weak Armor",
		Description: "If a physical attack hits this Pokemon, Defense is lowered by 1, Speed is raised by 1.",
	},
	"whitesmoke": {
		Name:        "White Smoke",
		Description: "Prevents other Pokemon from lowering this Pokemon's stat stages.",
	},
	"wonderguard": {
		Name:        "Wonder Guard",
		Description: "This Pokemon can only be damaged by supereffective moves and indirect damage.",
	},
	"wonderskin": {
		Name:        "Wonder Skin",
		Description: "Status moves with accuracy checks are 50% accurate when used on this Pokemon.",
	},
	"zenmode": {
		Name:        "Zen Mode",
		Description: "If Darmanitan, at end of turn changes Mode to Standard if > 1/2 max HP, else Zen.",
	},
	"mountaineer": {
		Name:        "Mountaineer",
		Description: "On switch-in, this Pokemon avoids all Rock-type attacks and Stealth Rock.",
	},
	"rebound": {
		Name:        "Rebound",
		Description: "On switch-in, blocks certain status moves and bounces them back to the user.",
	},
	"persistent": {
		Name:        "Persistent",
		Description: "The duration of certain field effects is increased by 2 turns if used by this Pokemon.",
	},
}
