<h1>data/enUS_Text/meta/StringList/Power_Necromancer_CorpseExplosion.stl</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/enUS_Text/meta/StringList/Power_Necromancer_CorpseExplosion.stl</td></tr><tr><td><b>Type</b></td><td>StringListDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1116113</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>arStrings</b></td><td><table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>name</code></td></tr><tr><td><b>szText</b></td><td><code>Corpse Explosion</code></td></tr><tr><td><b>hLabel</b></td><td><code>4062401</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>desc</code></td></tr><tr><td><b>szText</b></td><td><pre>{if:SF_10}{c_label}Essence Cost:{/c} {c_resource}{Resource Cost}{/c}
{/if}{if:ADVANCED_TOOLTIP}{c_label}Lucky Hit Chance: {/c}{c_resource}[{Combat_Effect_Chance_Script_Formula_Override:20}|%|]{/c}
{/if}{if:Mod.Miasma}Release a vile miasma from a Corpse, dealing {c_number}{dot:MIASMA_DOT_TOOLTIP}{/c} Shadow damage over {c_number}{buffduration:MIASMA_DOT_TOOLTIP}{/c} seconds.{else}Detonate a Corpse, dealing {c_number}{payload:DAMAGE}{/c} damage to surrounding enemies.{/if}</pre></td></tr><tr><td><b>hLabel</b></td><td><code>3707583</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>rankup_desc</code></td></tr><tr><td><b>szText</b></td><td><pre>
{icon:bullet,1.2} Damage {icon:arrow,1.2} {if:Mod.Miasma}{c_number}{dot:MIASMA_DOT_TOOLTIP}{/c}{else}{c_number}{payload:DAMAGE}{/c}{/if}</pre></td></tr><tr><td><b>hLabel</b></td><td><code>4074074671</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szText</b></td><td><code>{c_important}Corpse Explosion's{/c} radius is increased by {c_number}[{SF_1}*100|%|]{/c}.</code></td></tr><tr><td><b>hLabel</b></td><td><code>3591940403</code></td></tr><tr><td><b>szLabel</b></td><td><code>Mod0_Description</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szText</b></td><td><code>Enhanced Corpse Explosion</code></td></tr><tr><td><b>hLabel</b></td><td><code>2452871472</code></td></tr><tr><td><b>szLabel</b></td><td><code>Mod0_Name</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>Mod1_Description</code></td></tr><tr><td><b>szText</b></td><td><code>{c_important}Corpse Explosion{/c} becomes a {c_important}Darkness{/c} Skill and, instead of exploding, releases a vile miasma dealing {c_number}{dot:MIASMA_DOT_TOOLTIP}{/c} Shadow damage over {c_number}{SF_12}{/c} seconds.</code></td></tr><tr><td><b>hLabel</b></td><td><code>628601524</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>Mod1_Name</code></td></tr><tr><td><b>szText</b></td><td><code>Blighted Corpse Explosion</code></td></tr><tr><td><b>hLabel</b></td><td><code>2492006865</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>Mod2_Description</code></td></tr><tr><td><b>szText</b></td><td><code>{c_important}Corpse Explosion{/c} deals {c_number}[{SF_15}*100|%x|]{/c} increased damage to enemies that are Slowed, Stunned or {c_important}{u}Vulnerable{/u}{/c}. These damage bonuses can stack.</code></td></tr><tr><td><b>hLabel</b></td><td><code>1960229941</code></td></tr></table>


<table><tr><th colspan="100%">StringTableEntry</th></tr><tr><td><b>szLabel</b></td><td><code>Mod2_Name</code></td></tr><tr><td><b>szText</b></td><td><code>Plagued Corpse Explosion</code></td></tr><tr><td><b>hLabel</b></td><td><code>2531142258</code></td></tr></table>


</td></tr><tr><td><b>ptMapStringTable</b></td><td><code>0</code></td></tr></table>

