<h1>data/enUS_Text/meta/StringList/Power_Necromancer_Golem.stl</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/enUS_Text/meta/StringList/Power_Necromancer_Golem.stl</td></tr><tr><td><b>Type</b></td><td>StringListDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1115193</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>arStrings</b></td><td><table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>hLabel</b></td><td><code>4062401</code></td></tr><tr><td><b>szLabel</b></td><td><code>name</code></td></tr><tr><td><b>szText</b></td><td><code>Golem</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>desc</code></td></tr><tr><td><b>szText</b></td><td><pre>{if:SF_16}{c_label}Cooldown:{/c_label} {c_resource}{Cooldown Time}{/c_resource} seconds
{/if}{c_label}Passive:{/c} You are protected by a Golem with {c_number}{pet_health:PET_SPAWN_BLOOD}{/c} Life that attacks for {c_number}{payload:TOOLTIP_MELEE}{/c} damage. {if:SF_16}The Golem sheds Corpses as it takes damage.{/if}{if:SF_17}The Golem absorbs {c_number}[{SF_5}*100|%|]{/c} of damage you would take, and recovers Life when attacking.{/if}

{if:SF_16}{c_label}Active:{/c} Your Golem becomes {c_important}{u}Unstoppable{/u}{/c} and Taunts Nearby enemies and takes {c_number}[{SF_22}*100|%|]{/c} reduced damage for the next {c_number}6{/c} seconds.{/if}{if:SF_17}{c_label}Active:{/c} Your Golem becomes {c_important}{u}Unstoppable{/u}{/c} and drains the blood of enemies in the area, dealing {c_number}{payload:TOOLTIP_BLOOD_ACTIVE}{/c} damage and healing {c_number}[{SF_9}*100|%|]{/c} of its Life for each enemy drained. Damage and healing received are tripled if only one enemy is drained.{/if}{if:SF_18}{c_label}Active:{/c} Command your Golem to go to the targeted area, they become {c_important}{u}Unstoppable{/u}{/c} and slam their fists into the ground, dealing {c_number}{payload:TOOLTIP_SLAM}{/c} damage and Stunning surrounding enemies for {c_number}{SF_29}{/c} seconds.  This has a {c_number}{SF_20}{/c} second cooldown .{/if}

When your Golem dies, it respawns after {c_number}{SF_7}{/c} seconds.</pre></td></tr><tr><td><b>hLabel</b></td><td><code>3707583</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>rankup_desc</code></td></tr><tr><td><b>szText</b></td><td><pre>Life increased to {c_number}{pet_health:PET_SPAWN_BLOOD}{/c}. Attack damage increased to {c_number}{payload:TOOLTIP_MELEE}{/c}.

{if:SF_16}Active healing increased to {c_number}[{SF_14}*100|%|]{/c} per Corpse.{/if}{if:SF_17}Infusion damage increased to {c_number}{payload:TOOLTIP_BLOOD_ACTIVE}{/c}.{/if}</pre></td></tr><tr><td><b>hLabel</b></td><td><code>4074074671</code></td></tr></table>


</td></tr><tr><td><b>ptMapStringTable</b></td><td><code>0</code></td></tr></table>

