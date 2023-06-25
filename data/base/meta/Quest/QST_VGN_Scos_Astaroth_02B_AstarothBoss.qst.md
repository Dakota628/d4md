<h1>data/base/meta/Quest/QST_VGN_Scos_Astaroth_02B_AstarothBoss.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/QST_VGN_Scos_Astaroth_02B_AstarothBoss.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1042439</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>unk_942bcdb</b></td><td><code>1</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_af3a4c1</b></td><td></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>flParticipationRadius</b></td><td><code>3.75</code></td></tr><tr><td><b>eBountyTier</b></td><td><code>-1</code></td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>arQuestItems</b></td><td></td></tr><tr><td><b>ePlayerQuestType</b></td><td><code>-1</code></td></tr><tr><td><b>unk_c2e8448</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>arFollowers</b></td><td></td></tr><tr><td><b>unk_48a2b16</b></td><td><code>-1</code></td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>114</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	for i, idTrigger in ipairs(AREA_1_TRIGGERS) do
		Hydra.ActorDisable(idTrigger);
	end
	for i, idTrigger in ipairs(AREA_2_TRIGGERS) do
		Hydra.ActorDisable(idTrigger);
	end
end</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>2</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>179</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>BOSS</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3646615</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1426171770</code></td></tr><tr><td><b>eVariableType</b></td><td><code>2</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_7de5a3b2</th></tr><tr><td><b>dwType</b></td><td><code>2112201650</code></td></tr><tr><td><b>eParamType</b></td><td><code>3</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>flMinValue</b></td><td><code>0</code></td></tr><tr><td><b>flMaxValue</b></td><td><code>44.9900016784668</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>927407252</code></td></tr><tr><td><b>eEventType</b></td><td><code>7</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>180</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, flPrevHP, flHp)
	local idAmalgam = idActor
	
	if (Hydra.ActorGetSNOName(idAmalgam) ~= "astaroth_boss_mount") then
		local tBosses = Hydra.ActorsLinkedByActorInGroup(idSelf, "BOSS");
		
		for i, idBoss in ipairs(tBosses) do
			if (Hydra.ActorGetSNOName(idBoss) == "astaroth_boss_mount") then
				idAmalgam = idBoss;
				break;
			end
		end
	end
	
	local sDynamicMovementData = "AmalgamOfRage_Run";
	
	Hydra.QActorPower(idAmalgam, UNSTOPPABLE_POWER);	
	
	Hydra.QActorEnableAI(idAmalgam, false);
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "PreIntermission2LeapSpot")[1], sDynamicMovementData, false, false, false);
	
	local idSecondHouse = Hydra.ActorsLinkedByActorInGroup(idAmalgam, "SecondHouse")[1];
	
	if (Hydra.ActorIsValid(idSecondHouse)) then
		Hydra.QActorFaceActor(idAmalgam, idSecondHouse);
		Hydra.QActorPower(idAmalgam, SECOND_LEAP_HOUSE_POWER, idSecondHouse);
	else
		Hydra.EventSend("QST_Scos_Astaroth_02_AstarothEncounter_Phase2Complete");
	end
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "PreIntermission2MoveSpot")[1], sDynamicMovementData, false, false, false);

	Hydra.QActorPower(idAmalgam, LEAP_TO_NEXT_AREA_POWER, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "SecondJumpSpot")[1]);
	
	Hydra.QActorEnableAI(idAmalgam, true);
	
	Hydra.QActorPower(idAmalgam, UNSTOPPABLE_POWER);
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "Intermission2GoToSpot")[1], sDynamicMovementData, false, false, false);
end
</pre></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code>Hydra.QuestAdvancePhase()</code></td></tr><tr><td><b>dwUID</b></td><td><code>105</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>109</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>106</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3646615</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>BOSS</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2983118812</code></td></tr><tr><td><b>eVariableType</b></td><td><code>13</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamWorld</th></tr><tr><td><b>eParamType</b></td><td><code>25</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>dwType</b></td><td><code>1313165104</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>6</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamLevelArea</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoLevelArea</b></td><td><a href="..\LevelArea\Scos_Cerrigar_Boss.lvl">[DT_SNO] LevelArea: "Scos_Cerrigar_Boss"</a></td></tr><tr><td><b>dwType</b></td><td><code>1780051193</code></td></tr><tr><td><b>eParamType</b></td><td><code>5</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1584348101</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1243023061</code></td></tr><tr><td><b>eEventType</b></td><td><code>63</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sWorldName, sLevelAreaName)
	--Reset Boss
	local boss = Hydra.ActorsLinkedByActorInGroup(idSelf, "BOSS")
	for i, bossID in ipairs(boss) do
		Hydra.ActorFadeAndDelete(bossID, 2);
	end
end</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>111</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td><a href="#UKNOWN">[DT_SNO_NAME] LevelArea: "Scos_Cerrigar_Boss"</a>
</td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>110</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>0</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr></table>


</td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0.8999999761581421</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code>bPhase1Complete = false;</code></td></tr><tr><td><b>dwUID</b></td><td><code>104</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>5</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>107</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	Hydra.BossEncounterManuallyComplete(BOSS_LEVEL_AREA)
end
</pre></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>108</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>109</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>112</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td><a href="#UKNOWN">[DT_SNO_NAME] LevelArea: "Scos_Cerrigar_Boss"</a>
</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorPlayer, sLevelAreaName)
	if(Hydra.ActorIsAlive(idActorPlayer)) then
		--Boss Respawn Trigger Enabled
		Hydra.ActorEnable(AREA1_TRIGGER)
	end
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>113</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>snoCondition</b></td><td><a href="..\Condition\QST_Scos_Astaroth_02_Active.cnd">[DT_SNO] Condition: "QST_Scos_Astaroth_02_Active"</a></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>6</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamLevelArea</th></tr><tr><td><b>dwType</b></td><td><code>1780051193</code></td></tr><tr><td><b>eParamType</b></td><td><code>5</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoLevelArea</b></td><td><a href="..\LevelArea\Scos_Cerrigar_Boss.lvl">[DT_SNO] LevelArea: "Scos_Cerrigar_Boss"</a></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1584348101</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>143506636</code></td></tr><tr><td><b>eEventType</b></td><td><code>44</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>158</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>6</code></td></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Astaroth_Boss.mrk">[DT_SNO] MarkerSet: "Astaroth_Boss"</a></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>159</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
	Hydra.EventSend("QST_VGN_Scos_Astaroth_02B_AstarothBoss_Start");
end
</pre></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>120</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>187</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)
		
	Hydra.WaitForConversationFinished(IGC_SLA);
	
	--Manually Start Encounter
	Hydra.BossEncounterManuallyQueueStart(BOSS_LEVEL_AREA)
	
	local tBosses = Hydra.ActorsEnableLinkedByActorInGroup(idSelf, "BOSS")
	
	local tNPCs = Hydra.ActorsEnableLinkedByActorInGroup(idSelf, "OnFireNPC");
	
	local S_FEMALE_ACTOR_NAME = Hydra.ReferenceActor("astaroth_boss_villager_female_onfire");

	if (#tNPCs > 0) then
		local snoStaggerPower = Hydra.PowerSNOFromName(Hydra.ReferencePower("astaroth_boss_mount_stagger"));
		
		for i, idNPC in ipairs(tNPCs) do
			
			local sActorName = Hydra.ActorGetSNOName(idNPC);		
			
			if (sActorName == S_FEMALE_ACTOR_NAME) then
				Hydra.BuffApplyToTarget(snoStaggerPower, "NPC_ON_FIRE_F", idNPC);
			else
				Hydra.BuffApplyToTarget(snoStaggerPower, "NPC_ON_FIRE_M", idNPC);
			end
			
			for b, idBoss in ipairs(tBosses) do 
				Hydra.ActorAIForcePrimaryTarget(idNPC, idBoss);
			end
			
		end
	end

	--Disable Trigger
	Hydra.ActorDisable(AREA1_TRIGGER)
	
	-- Disable Traversals so players can't teleport through stuff
	local tTraversals = Hydra.ActorsLinkedByActorInGroup(idSelf, "TraversalLocator");
	
	if (#tTraversals > 0) then
		for i, idTraversal in ipairs(tTraversals) do
			SetTraversalEnabled(idTraversal, false)
		end
	end
end</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>121</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>3926827530</code></td></tr><tr><td><b>szString</b></td><td><code>QST_VGN_Scos_Astaroth_02B_AstarothBoss_Start</code></td></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr></table>


</td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>114</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code>Hydra.QuestAdvancePhase()</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>185</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>104</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3646615</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>BOSS</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1426171770</code></td></tr><tr><td><b>eVariableType</b></td><td><code>2</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_7de5a3b2</th></tr><tr><td><b>dwType</b></td><td><code>2112201650</code></td></tr><tr><td><b>eParamType</b></td><td><code>3</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>flMinValue</b></td><td><code>0</code></td></tr><tr><td><b>flMaxValue</b></td><td><code>74.98999786376953</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>927407252</code></td></tr><tr><td><b>eEventType</b></td><td><code>7</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>188</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, flPrevHP, flHp)
	local idAmalgam = idActor
	
	if (Hydra.ActorGetSNOName(idAmalgam) ~= "astaroth_boss_mount") then
		local tBosses = Hydra.ActorsLinkedByActorInGroup(idSelf, "BOSS");
		
		for i, idBoss in ipairs(tBosses) do
			if (Hydra.ActorGetSNOName(idBoss) == "astaroth_boss_mount") then
				idAmalgam = idBoss;
				break;
			end
		end
	end
	
	local sDynamicMovementData = "AmalgamOfRage_Run";
	
	Hydra.QActorPower(idAmalgam, UNSTOPPABLE_POWER);
	
	Hydra.QActorEnableAI(idAmalgam, false);
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "PreIntermission1LeapSpot")[1], sDynamicMovementData, false, false, false);
	
	local idFirstHouse = Hydra.ActorsLinkedByActorInGroup(idAmalgam, "FirstHouse")[1];
	
	if (Hydra.ActorIsValid(idFirstHouse)) then
		Hydra.QActorFaceActor(idAmalgam, idFirstHouse);
		Hydra.QActorPower(idAmalgam, FIRST_LEAP_HOUSE_POWER, idFirstHouse);
	else
		Hydra.EventSend("QST_Scos_Astaroth_02_AstarothEncounter_Phase1Complete");
	end
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "PreIntermission1MoveSpot")[1], sDynamicMovementData, false, false, false);

	Hydra.QActorPower(idAmalgam, LEAP_TO_NEXT_AREA_POWER, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "FirstJumpSpot")[1]);
	
	Hydra.QActorEnableAI(idAmalgam, true);
	
	Hydra.QActorPower(idAmalgam, UNSTOPPABLE_POWER);
	
	Hydra.QActorMoveToLocationDynamicSpeed(idAmalgam, Hydra.ActorsLinkedByActorInGroup(idAmalgam, "Intermission1GoToSpot")[1], sDynamicMovementData, false, false, false);
end
</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>191</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>0</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2983118812</code></td></tr><tr><td><b>eVariableType</b></td><td><code>13</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamWorld</th></tr><tr><td><b>dwType</b></td><td><code>1313165104</code></td></tr><tr><td><b>eParamType</b></td><td><code>25</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamLevelArea</th></tr><tr><td><b>eParamType</b></td><td><code>5</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoLevelArea</b></td><td><a href="..\LevelArea\Scos_Cerrigar_Boss.lvl">[DT_SNO] LevelArea: "Scos_Cerrigar_Boss"</a></td></tr><tr><td><b>dwType</b></td><td><code>1780051193</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1584348101</code></td></tr><tr><td><b>eVariableType</b></td><td><code>6</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1243023061</code></td></tr><tr><td><b>eEventType</b></td><td><code>63</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sWorldName, sLevelAreaName)
	--Reset Boss
	local boss = Hydra.ActorsLinkedByActorInGroup(idSelf, "BOSS")
	for i, bossID in ipairs(boss) do
		Hydra.ActorFadeAndDelete(bossID, 2);
	end
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td><a href="#UKNOWN">[DT_SNO_NAME] LevelArea: "Scos_Cerrigar_Boss"</a>
</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>192</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwUID</b></td><td><code>187</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>206</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>204</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)
	for i, idTrigger in ipairs(AREA_1_TRIGGERS) do
		Hydra.ActorEnable(idTrigger);
	end
end
</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>205</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>886169572</code></td></tr><tr><td><b>szString</b></td><td><code>QST_VGN_Scos_Astaroth_02B_AstarothBoss_AstarothInArea2</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr></table>

</td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>207</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)
	for i, idTrigger in ipairs(AREA_2_TRIGGERS) do
		Hydra.ActorEnable(idTrigger);
	end
end
</pre></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>208</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>szString</b></td><td><code>QST_VGN_Scos_Astaroth_02B_AstarothBoss_AstarothInArea3</code></td></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>886169573</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>209</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>210</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1637442350</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>48</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(listValidActors, idActorDetector)
	local idAmalgam = Hydra.ActorsLinkedByActorInGroup(idActorDetector, "AMALGAM")[1];
	
	for i, idPlayer in ipairs(listValidActors) do
		if Hydra.ActorIsAlive(idPlayer) then
			Hydra.ActorTriggerPowerAtTarget(idAmalgam, idPlayer, METEOR_DROP_POWER, true);
		else
			-- Determine what area we are in
			if (Hydra.ActorIsEnabled(Hydra.ActorsInGroup("Astaroth_Area_2")[1])) then
				-- Port to Area 3 Gravespots
				local tGraveSpots = Hydra.ActorsEnableLinkedByActorInGroup(idActorDetector, "AREA_3_GRAVESPOT");
				
				if (#tGraveSpots > 0) then
					Hydra.ActorWarpToActor(idPlayer, tGraveSpots[Hydra.GetRandomInt(1, #tGraveSpots)]);
				end
			else
				-- Port to Area 2 Gravespots
				local tGraveSpots = Hydra.ActorsEnableLinkedByActorInGroup(idActorDetector, "AREA_2_GRAVESPOT");
				
				if (#tGraveSpots > 0) then
					Hydra.ActorWarpToActor(idPlayer, tGraveSpots[Hydra.GetRandomInt(1, #tGraveSpots)]);
				end
			end
		end
	end
end
</pre></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>212</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(listValidActors, idActorDetector)
	local idAmalgam = Hydra.ActorsLinkedByActorInGroup(idActorDetector, "AMALGAM")[1];
	
	for i, idPlayer in ipairs(listValidActors) do
		if Hydra.ActorIsAlive(idPlayer) then
			Hydra.ActorTriggerPowerAtTarget(idAmalgam, idPlayer, METEOR_DROP_POWER, true);
		else
			-- Port to Area 3 Gravespots
			local tGraveSpots = Hydra.ActorsEnableLinkedByActorInGroup(idActorDetector, "AREA_3_GRAVESPOT");
			
			if (#tGraveSpots > 0) then
				Hydra.ActorWarpToActor(idPlayer, tGraveSpots[Hydra.GetRandomInt(1, #tGraveSpots)]);
			end
		end
	end
end
</pre></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>214</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1637442351</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>48</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>eType</b></td><td><code>2</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr></table>


</td></tr><tr><td><b>eInstanceQuestType</b></td><td><code>0</code></td></tr><tr><td><b>eVignetteType</b></td><td><code>0</code></td></tr><tr><td><b>szOnStartupScript</b></td><td><pre>S_LEVEL_AREA = Hydra.ActorGetLevelArea(idSelf)
AREA1_TRIGGER = Hydra.ActorsLinkedByActorInGroup(idSelf, "TRIGGER")[1];</pre></td></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr><tr><td><b>szUserFunctionsScript</b></td><td><code></code></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_c18cabd</b></td><td><code>0</code></td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>arRequiredReputations</b></td><td></td></tr><tr><td><b>eQuestType</b></td><td><code>1</code></td></tr><tr><td><b>eEventQuestType</b></td><td><code>0</code></td></tr><tr><td><b>szOnShutdownScript</b></td><td><code></code></td></tr><tr><td><b>unk_d2181f0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_14dee1b</b></td><td><code>1</code></td></tr><tr><td><b>unk_8881b0e</b></td><td><code>20</code></td></tr><tr><td><b>dwNextUID</b></td><td><code>235</code></td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>unk_b83e7b1</b></td><td><code>0</code></td></tr><tr><td><b>eRepeatType</b></td><td><code>0</code></td></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>unk_2aa5f20</b></td><td></td></tr><tr><td><b>unk_b43b442</b></td><td></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr></table>

