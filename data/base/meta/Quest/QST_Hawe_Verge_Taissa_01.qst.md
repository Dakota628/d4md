<h1>data/base/meta/Quest/QST_Hawe_Verge_Taissa_01.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/QST_Hawe_Verge_Taissa_01.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1188154</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>unk_2aa5f20</b></td><td><a href="QST_Hawe_Verge_Taissa_01_Cellar.qst.md">[DT_SNO] Quest: "QST_Hawe_Verge_Taissa_01_Cellar"</a>
</td></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwNextUID</b></td><td><code>133</code></td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szUserFunctionsScript</b></td><td><code></code></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(-871.857971, 463.657990, 21.287500)</td></tr><tr><td><b>snoLevelArea</b></td><td><a href="#UKNOWN">[DT_SNO] LevelArea: "Hawe_Verge_RuinedTower"</a></td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>unk_14dee1b</b></td><td><code>0</code></td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>87</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	
	idOrinaClone = Hydra.CreatePersonalClone(idActorOrina, true)
	if Hydra.ActorIsValid(idOrinaClone) then
    	Hydra.ActorEnableAI(idOrinaClone, false)
    	Hydra.QActorMoveToLocationDynamicSpeed(idOrinaClone, locTowerExit, "NPC_WALK", true)
		Hydra.QActorDeletePersonalClone(idOrinaClone, 0.3)
	end
	if (Hydra.QuestIsRestoredFromSave(false)) then
		Hydra.EventSendToActor("OrinaLeaves", idSelf, idSelf)
	end
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>86</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>120</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="#UKNOWN">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>3926636179</code></td></tr><tr><td><b>eEventType</b></td><td><code>27</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamSelf</th></tr><tr><td><b>dwType</b></td><td><code>1210992754</code></td></tr><tr><td><b>eParamType</b></td><td><code>11</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="#UKNOWN">[DT_SNO] MarkerSet: "Hawe_Verge (SideQuest_Taissa)"</a></td></tr><tr><td><b>nID</b></td><td><code>4</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1072510348</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>121</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorPlayer, idActorGizmo)
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--Hydra.QuestAdvancePhase()
</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>122</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>snoWorld</b></td><td><a href="#UKNOWN">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorPlayer, idActorGizmo)
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>123</code></td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType