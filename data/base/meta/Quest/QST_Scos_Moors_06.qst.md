<h1>data/base/meta/Quest/QST_Scos_Moors_06.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/QST_Scos_Moors_06.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>875681</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>unk_af3a4c1</b></td><td></td></tr><tr><td><b>unk_d2181f0</b></td><td><code>11341</code></td></tr><tr><td><b>unk_8881b0e</b></td><td><code>20</code></td></tr><tr><td><b>dwNextUID</b></td><td><code>28</code></td></tr><tr><td><b>unk_942bcdb</b></td><td><code>1</code></td></tr><tr><td><b>arRequiredReputations</b></td><td></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>gbidSurveyType</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3340288678</code></td></tr></table>

</td></tr><tr><td><b>eRepeatType</b></td><td><code>0</code></td></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(-2280.780029, -131.085999, 11.727000)</td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>szOnStartupScript</b></td><td><pre>

-- Set groups of Tendrils

tGrowthFamily01 = Hydra.ActorsLinkedByActor(ID_GROWTH_01)
tGrowthFamily02 = Hydra.ActorsLinkedByActor(ID_GROWTH_02)
tGrowthFamily05 = Hydra.ActorsLinkedByActor(ID_GROWTH_05)

--Make Profane Clot Invulnerable and Untargetable
Hydra.ActorSetInvulnerable(ID_CLOT, true)
Hydra.ActorSetTargetable(ID_CLOT, false)

-- hiding any wolf already in scene
tWolves = Hydra.ActorsNearActor(ID_NAFAIN, 8.0, "NPC_QST_Scos_NafainWolf_No_Icon")
for i = 1, #tWolves do
	if Hydra.ActorIsValid(tWolves[i]) then
		if Hydra.ActorIsEnabled(tWolves[i]) then
			Hydra.ActorDisable(tWolves[i])
		end
	end
end

</pre></td></tr><tr><td><b>eQuestType</b></td><td><code>2</code></td></tr><tr><td><b>unk_48a2b16</b></td><td><code>-1</code></td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>szUserFunctionsScript</b></td><td><code></code></td></tr><tr><td><b>eVignetteType</b></td><td><code>0</code></td></tr><tr><td><b>snoLevelArea</b></td><td><a href="..\LevelArea\Scos_Moors_BloodVale.lvl">[DT_SNO] LevelArea: "Scos_Moors_BloodVale"</a></td></tr><tr><td><b>unk_14dee1b</b></td><td><code>0</code></td></tr><tr><td><b>unk_b43b442</b></td><td></td></tr><tr><td><b>eInstanceQuestType</b></td><td><code>2</code></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>unk_c18cabd</b></td><td><code>0</code></td></tr><tr><td><b>eBountyTier</b></td><td><code>-1</code></td></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr><tr><td><b>eEventQuestType</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a4ec7f</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>flParticipationRadius</b></td><td><code>12</code></td></tr><tr><td><b>unk_b83e7b1</b></td><td><code>3</code></td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>arFollowers</b></td><td><table><tr><th colspan="100%">QuestFollower</th></tr><tr><td><b>snoActor</b></td><td><a href="..\Actor\nafain_wolf_follower.acr">[DT_SNO] Actor: "nafain_wolf_follower"</a></td></tr><tr><td><b>unk_74ee52d</b></td><td><table><tr><th colspan="100%">Type_da8a2315</th></tr><tr><td><b>unk_6c68319</b></td><td><code>3</code></td></tr><tr><td><b>unk_d64b2c9</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_Nafain_Base).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_Nafain_Base)"</a></td></tr><tr><td><b>nID</b></td><td><code>102</code></td></tr></table>

</td></tr><tr><td><b>eVisibility</b></td><td><code>0</code></td></tr><tr><td><b>unk_209d19e</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arPhases</b></td><td><table><tr><th colspan="100%">Type_bc83f219</th></tr><tr><td><b>dwPhase</b></td><td><code>13</code></td></tr><tr><td><b>unk_74ee52d</b></td><td><table><tr><th colspan="100%">Type_da8a2315</th></tr><tr><td><b>unk_6c68319</b></td><td><code>3</code></td></tr><tr><td><b>unk_d64b2c9</b></td><td><code>1</code></td></tr><tr><td><b>snoSummonPower</b></td><td><a href="..\Power\Druid_Wolf_Howl.pow">[DT_SNO] Power: "Druid_Wolf_Howl"</a></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_Nafain_Base).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_Nafain_Base)"</a></td></tr><tr><td><b>nID</b></td><td><code>102</code></td></tr></table>

</td></tr><tr><td><b>eVisibility</b></td><td><code>0</code></td></tr><tr><td><b>unk_209d19e</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr></table>


<table><tr><th colspan="100%">Type_bc83f219</th></tr><tr><td><b>unk_74ee52d</b></td><td><table><tr><th colspan="100%">Type_da8a2315</th></tr><tr><td><b>unk_6c68319</b></td><td><code>0</code></td></tr><tr><td><b>unk_d64b2c9</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>eVisibility</b></td><td><code>0</code></td></tr><tr><td><b>unk_209d19e</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>dwPhase</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">Type_bc83f219</th></tr><tr><td><b>dwPhase</b></td><td><code>17</code></td></tr><tr><td><b>unk_74ee52d</b></td><td><table><tr><th colspan="100%">Type_da8a2315</th></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>eVisibility</b></td><td><code>0</code></td></tr><tr><td><b>unk_209d19e</b></td><td><code>0</code></td></tr><tr><td><b>unk_6c68319</b></td><td><code>0</code></td></tr><tr><td><b>unk_d64b2c9</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>szOnShutdownScript</b></td><td><code></code></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>unk_c2e8448</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>-- Swaps out growths for ones that can be targeted and destroyed
for i = 1, #T_GRP_GROWTHS_SERVER_PROPS do

	if Hydra.ActorIsValid(T_GRP_GROWTHS_SERVER_PROPS[i]) then
	
		Hydra.ActorDelete(T_GRP_GROWTHS_SERVER_PROPS[i])
		Hydra.ActorEnable(T_GRP_GROWTHS[i])
	end

end</pre></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(-2176.179932, -203.526993, 2.266440)</td></tr><tr><td><b>unk_ed8215b</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>1</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arLevelAreas</b></td><td><a href="..\LevelArea\Scos_Moors_BloodVale.lvl">[DT_SNO] LevelArea: "Scos_Moors_BloodVale"</a>
</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>10</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>2</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)

	-- Step through the linked list of tendrils and fade them out

	for i = 1, #tGrowthFamily01 do
		if Hydra.ActorIsValid(tGrowthFamily01[i]) then
			flRandNum = Hydra.GetRandomFloat(2, 2.5)
			if Hydra.ActorIsGizmo(tGrowthFamily01[i]) then
				Hydra.ActorGizmoSetDisabled(tGrowthFamily01[i], false)
				Hydra.ActorKill(tGrowthFamily01[i])
			else
				Hydra.ActorFadeAndDelete(tGrowthFamily01[i], flRandNum)
			end
		end
	end 
	
end
</pre></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwUID</b></td><td><code>8</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)

	-- Step through the linked list of tendrils and fade them out

	for i = 1, #tGrowthFamily02 do
		if Hydra.ActorIsValid(tGrowthFamily02[i]) then
			flRandNum = Hydra.GetRandomFloat(2, 2.5)
			if Hydra.ActorIsGizmo(tGrowthFamily02[i]) then
				Hydra.ActorGizmoSetDisabled(tGrowthFamily02[i], false)
				Hydra.ActorKill(tGrowthFamily02[i])
			else
				Hydra.ActorFadeAndDelete(tGrowthFamily02[i], flRandNum)
			end
		end
	end 

end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>14</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>16</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)

	for i = 1, #tGrowthFamily05 do
		if Hydra.ActorIsValid(tGrowthFamily05[i]) then
			flRandNum = Hydra.GetRandomFloat(2, 2.5)
			if Hydra.ActorIsGizmo(tGrowthFamily05[i]) then
				Hydra.ActorGizmoSetDisabled(tGrowthFamily05[i], false)
				Hydra.ActorKill(tGrowthFamily05[i])
			else
				Hydra.ActorFadeAndDelete(tGrowthFamily05[i], flRandNum)
			end
		end
	end 

end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>203</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>
</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>9</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>17</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)


end
</pre></td></tr><tr><td><b>unk_6a71535</b></td><td><code>1</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>3</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>10</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td>Vector(-2116.000000, -170.000000)
Vector(-2116.000000, -205.000000)
Vector(-2140.000000, -268.000000)
Vector(-2200.000000, -278.000000)
Vector(-2208.000000, -278.000000)
Vector(-2260.000000, -264.000000)
Vector(-2264.000000, -260.000000)
Vector(-2264.000000, -244.000000)
Vector(-2260.000000, -110.000000)
Vector(-2258.000000, -108.000000)
Vector(-2244.000000, -106.000000)
Vector(-2192.000000, -102.000000)
Vector(-2190.000000, -102.000000)
Vector(-2184.000000, -104.000000)
Vector(-2180.000000, -106.000000)
Vector(-2166.000000, -116.000000)
Vector(-2146.000000, -134.000000)
Vector(-2136.000000, -144.000000)
Vector(-2118.000000, -166.000000)
</td></tr><tr><td><b>vCenterPos</b></td><td>Vector(-2190.000000, -190.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>106.52699279785156</code></td></tr><tr><td><b>bValid</b></td><td><code>1</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>1</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td>Vector(-2123.969971, -227.953003)
Vector(-2142.330078, -263.670013)
Vector(-2212.719971, -275.841003)
Vector(-2255.439941, -259.152008)
Vector(-2258.870117, -226.238998)
Vector(-2220.050049, -178.617996)
Vector(-2181.000000, -175.479996)
Vector(-2141.699951, -199.914993)
</td></tr><tr><td><b>vCenterPos</b></td><td>Vector(-2191.419922, -225.660492)</td></tr><tr><td><b>flRadius</b></td><td><code>72.25125885009766</code></td></tr><tr><td><b>bValid</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td><a href="..\LevelArea\Scos_Moors_BloodVale.lvl">[DT_SNO] LevelArea: "Scos_Moors_BloodVale"</a>
</td></tr><tr><td><b>dwFlags</b></td><td><code>544</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamActorGroup</th></tr><tr><td><b>gbidActorGroup</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>3830087414</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3481819086</code></td></tr><tr><td><b>eParamType</b></td><td><code>35</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>-- Deletes blockage objects from the doorway

for i = 1, #T_GRP_DOOR_TENDRILS do
	if Hydra.ActorIsValid(T_GRP_DOOR_TENDRILS[i]) then
		flRandNum = Hydra.GetRandomFloat(1, 2.5)
		Hydra.ActorFadeAndDelete(T_GRP_DOOR_TENDRILS[i], flRandNum)
	end
end

--Make Profane Clot Vulnerable
Hydra.ActorSetInvulnerable(ID_CLOT, false)
Hydra.ActorSetTargetable(ID_CLOT, true)

Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>331</code></td></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Game).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Game)"</a></td></tr></table>

</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>252</code></td></tr></table>

</td></tr><tr><td><b>unk_ed8215b</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(-2247.275879, -245.125885, 2.317868)</td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>5</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>1</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>4</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>4</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>3</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>4</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(-2246.090088, -244.886993, 2.958460)</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>13</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>199</code></td></tr></table>

</td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_ed8215b</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>11</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>0</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>12</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)

	Hydra.ActorPlayConversation(S_CONVO_NO_ENTRY, idActor)

end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>199</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>14</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwUID</b></td><td><code>15</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2983118812</code></td></tr><tr><td><b>eVariableType</b></td><td><code>13</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamWorld</th></tr><tr><td><b>dwType</b></td><td><code>1313165104</code></td></tr><tr><td><b>eParamType</b></td><td><code>25</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\CSD_Scos_Moors_HeartoftheMoors.wrl">[DT_SNO] World: "CSD_Scos_Moors_HeartoftheMoors"</a></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1542277347</code></td></tr><tr><td><b>eEventType</b></td><td><code>43</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorPlayer, sWorldName)
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>
</pre></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>
</pre></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>26</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>dwUID</b></td><td><code>27</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	Hydra.ActorFadeAndDelete(ID_WHITE_WOLF, 1)
end
</pre></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwUID</b></td><td><code>17</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>-- Deletes blockage objects from the doorway

-- Disabled these in axe and commenting them out here for F&F.  The duplicate collision was preventing players from directly damaging the clot.  -RMaille 06/28/22
--Hydra.ActorDelete(ID_DOOR_BLOCKAGE)

Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>24</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>5</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwUID</b></td><td><code>25</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)

--Play gore explosion
Hydra.ActorPlayEffectGroup(idActorKiller, "fxKit_aoe_explosion_cylindrical_necroBlood")
Hydra.ActorPlaySound(idActorKiller,"Play_MON_Fallen_Lunatic_FOLY_Death_Gore_Explosion_3P")

--Hydra.Wait(1.5)

end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>flRadius</b></td><td><code>12</code></td></tr><tr><td><b>bValid</b></td><td><code>1</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>1</code></td></tr><tr><td><b>vPolygonPoints</b></td><td>Vector(-2247.275879, -245.125885)
</td></tr><tr><td><b>vCenterPos</b></td><td>Vector(-2247.275879, -245.125885)</td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>1</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>snoWorld</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>arRegionOverrides</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(-2191.419922, -225.660492)</td></tr><tr><td><b>flRadius</b></td><td><code>72.25125885009766</code></td></tr><tr><td><b>bValid</b></td><td><code>1</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>1</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td>Vector(-2123.969971, -227.953003)
Vector(-2142.330078, -263.670013)
Vector(-2212.719971, -275.841003)
Vector(-2255.439941, -259.152008)
Vector(-2258.870117, -226.238998)
Vector(-2220.050049, -178.617996)
Vector(-2181.000000, -175.479996)
Vector(-2141.699951, -199.914993)
</td></tr></table>


</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamMarkerHandle</th></tr><tr><td><b>dwType</b></td><td><code>3311139638</code></td></tr><tr><td><b>eParamType</b></td><td><code>33</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>tMarkerHandle</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>252</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td><a href="..\LevelArea\Scos_Moors_BloodVale.lvl">[DT_SNO] LevelArea: "Scos_Moors_BloodVale"</a>
</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Scos_Moors_BloodVale_Campaign).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Scos_Moors_BloodVale_Campaign)"</a></td></tr><tr><td><b>nID</b></td><td><code>252</code></td></tr></table>

</td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(-2247.275879, -245.125885, 2.317868)</td></tr><tr><td><b>unk_ed8215b</b></td><td><a href="..\World\Sanctuary_Eastern_Continent.wrl">[DT_SNO] World: "Sanctuary_Eastern_Continent"</a></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>arQuestItems</b></td><td></td></tr><tr><td><b>unk_2aa5f20</b></td><td></td></tr><tr><td><b>ePlayerQuestType</b></td><td><code>0</code></td></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>snoMarkerSet</b></td><td><a href="..\MarkerSet\Scos_Moors (Game).mrk">[DT_SNO] MarkerSet: "Scos_Moors (Game)"</a></td></tr><tr><td><b>nID</b></td><td><code>17</code></td></tr></table>

</td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr></table>

