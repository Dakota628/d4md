<h1>data/base/meta/Quest/QST_VGN_Step_FuryAgainstFate_Holdout.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/QST_VGN_Step_FuryAgainstFate_Holdout.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1275164</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>flParticipationRadius</b></td><td><code>30</code></td></tr><tr><td><b>unk_b83e7b1</b></td><td><code>0</code></td></tr><tr><td><b>unk_942bcdb</b></td><td><code>1</code></td></tr><tr><td><b>arRequiredReputations</b></td><td></td></tr><tr><td><b>unk_2aa5f20</b></td><td></td></tr><tr><td><b>eInstanceQuestType</b></td><td><code>1</code></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_14dee1b</b></td><td><code>0</code></td></tr><tr><td><b>dwNextUID</b></td><td><code>48</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwUID</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>6</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)

--Cannibal Cry
Hydra.EventSend("QST_FuryAgainstFate_CannibalWalla")
--Activate Wave 1 Spawners
	Hydra.Wait(2)
	
	for i, idSpawner in ipairs(T_WAVE01_SPAWNERS) do
		Hydra.ActorEnable(idSpawner)
	end
	
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>2</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>1444490448</code></td></tr><tr><td><b>szString</b></td><td><code>QST_FuryAgainstFate_HoldoutStart</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>1</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nNeeded</b></td><td><code>10</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>37</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367108</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr></table>


</td></tr></table>


</td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>36</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--Disabling Wave 01 Spawners, Variable reset for nil safety
tWave01Spawners = Hydra.ActorsInGroup("QST_Step_South_FuryAgainstFate_Wave01Spawners")

for i, idSpawner in ipairs(tWave01Spawners) do

	Hydra.ActorDisable(idSpawner)

end</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, szGroup)

--Activate Wave 2 Spawners
	Hydra.Wait(4)
	
	for i, idSpawner in ipairs(T_WAVE02_SPAWNERS) do
		Hydra.ActorEnable(idSpawner)
	end
	
--Cannibal Cry
Hydra.EventSend("QST_FuryAgainstFate_CannibalWalla")
	
--Ealda Insult Event Cue 01
Hydra.EventSend("QST_FuryAgainstFate_EaldaInsult01")
	
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>39</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3095850554</code></td></tr><tr><td><b>eVariableType</b></td><td><code>12</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367108</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3682775607</code></td></tr><tr><td><b>eEventType</b></td><td><code>39</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>38</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--Disabling Wave 01 Spawners, Variable reset for nil safety
tWave02Spawners = Hydra.ActorsInGroup("QST_Step_South_FuryAgainstFate_Wave02Spawners")

for i, idSpawner in ipairs(tWave02Spawners) do

	Hydra.ActorDisable(idSpawner)

end
</pre></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>40</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>41</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367109</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>10</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>42</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, szGroup)

--Activate Wave 3 Spawners
	Hydra.Wait(4)
	
	for i, idSpawner in ipairs(T_WAVE03_SPAWNERS) do
		Hydra.ActorEnable(idSpawner)
	end
	
--Cannibal Cry
Hydra.EventSend("QST_FuryAgainstFate_CannibalWalla")

--Ealda Insult Event Cue 02
Hydra.EventSend("QST_FuryAgainstFate_EaldaInsult02")
	
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3095850554</code></td></tr><tr><td><b>eVariableType</b></td><td><code>12</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367109</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3682775607</code></td></tr><tr><td><b>eEventType</b></td><td><code>39</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>43</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>44</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--Disabling Wave 01 Spawners, Variable reset for nil safety
tWave03Spawners = Hydra.ActorsInGroup("QST_Step_South_FuryAgainstFate_Wave03Spawners")

for i, idSpawner in ipairs(tWave03Spawners) do

	Hydra.ActorDisable(idSpawner)

end</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>45</code></td></tr><tr><td><b>nNeeded</b></td><td><code>10</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367110</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr></table>


</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>46</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>39</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3095850554</code></td></tr><tr><td><b>eVariableType</b></td><td><code>12</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3382367110</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3682775607</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, szGroup)

--Despawn Prefab
	Hydra.ActorFlagDynamicPrefabGroupForDespawn(idSelf)
	
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>47</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_b43b442</b></td><td></td></tr><tr><td><b>unk_af3a4c1</b></td><td></td></tr><tr><td><b>eRepeatType</b></td><td><code>0</code></td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>szUserFunctionsScript</b></td><td><code></code></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>arFollowers</b></td><td></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>unk_48a2b16</b></td><td><code>-1</code></td></tr><tr><td><b>unk_d2181f0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szOnStartupScript</b></td><td><code></code></td></tr><tr><td><b>szOnShutdownScript</b></td><td><code></code></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>arQuestItems</b></td><td></td></tr><tr><td><b>ePlayerQuestType</b></td><td><code>0</code></td></tr><tr><td><b>eEventQuestType</b></td><td><code>0</code></td></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>unk_c2e8448</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>eVignetteType</b></td><td><code>3</code></td></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>eBountyTier</b></td><td><code>-1</code></td></tr><tr><td><b>eQuestType</b></td><td><code>1</code></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_8881b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_c18cabd</b></td><td><code>0</code></td></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr></table>

<h2>Quest Details (WIP)</h2><h3>Phase Order</h3>