<h1>data/base/meta/Quest/LE_Ambush_Standard.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/LE_Ambush_Standard.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>952963</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>eVignetteType</b></td><td><code>0</code></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>szOnStartupScript</b></td><td><pre>-- ======= REFERENCES =======	
	S_KEY_HIDE = "hide"
	S_KEY_UNHIDE = "unhide"
	S_KEY_GETUP = "getup"
	
	T_TARGETABLE = {}
	T_TARGETABLE[S_KEY_HIDE] = false
	T_TARGETABLE[S_KEY_UNHIDE] = true
	T_TARGETABLE[S_KEY_GETUP] = true
	
	T_INTRO_ANIMKEYS = {}
	T_INTRO_ANIMKEYS[S_KEY_HIDE] = Hydra.ReferencePower("AnimKey_OnChest_Hide_Intro")
	T_INTRO_ANIMKEYS[S_KEY_UNHIDE] = Hydra.ReferencePower("AnimKey_OnChest_Hide_Outro")
	T_INTRO_ANIMKEYS[S_KEY_GETUP] = Hydra.ReferencePower("AnimKey_OnChest_Getup")

	T_IDLE_ANIMKEYS = {}
	T_IDLE_ANIMKEYS[S_KEY_HIDE] = Hydra.ReferencePower("AnimKey_OnChest_Hide_Idle")
	T_IDLE_ANIMKEYS[S_KEY_UNHIDE] = Hydra.ReferencePower("AnimKey_Neutral")
	
	-- making getup idle anim a random wounded anim
	T_END_ANIMKEYS = {Hydra.ReferencePower("AnimKey_Wounded_Standing_Idle"), Hydra.ReferencePower("AnimKey_Wounded_Standing_Idle_Arm"), Hydra.ReferencePower("AnimKey_Idle_Sad")}
	T_IDLE_ANIMKEYS[S_KEY_GETUP] = Hydra.ReferencePower(T_END_ANIMKEYS[Hydra.GetRandomInt(1, #T_END_ANIMKEYS)])
	
	T_NPC_END_ACTIONS = {"idle", "idle", "cry", "cry", "pray", "pray"}
	
	T_NPC_DEATH_FX = {Hydra.ReferenceEffectGroup("fxKit_blood_splash_a"), Hydra.ReferenceEffectGroup("fxKit_blood_splash_b"), Hydra.ReferenceEffectGroup("fxKit_blood_splash_c"), Hydra.ReferenceEffectGroup("fxKit_blood_spurt_b")}
		
-- ======= VARIABLES =======
--	Constants
	ID_THIS_QUEST = Hydra.QuestGetOwner()
	N_RANGE_QUEST = 16
	
	--initialize wagon data structures
	T_VILLAGERS = {}
	
	for i, idWagon in pairs(T_WAGONS) do
		T_VILLAGERS[idWagon] = Hydra.ActorsLinkedByActorInGroup(idWagon, "NPC")[1]
	end
	
	playVillagerAnimAndIdle(S_KEY_HIDE)
	
	for i = 1, #T_CORPSES do
		Hydra.ActorGizmoSetDisabled(T_CORPSES[i], true)
	end 
	

	
	FYShuffle(T_WAGONS)
	
--	local nWagonsToDisable = -- world tier implementation
--	{
--		[0] = 2,
--		[1] = 2,
--		[2] = 1,
--		[3] = 0,
--		[4] = 0
--	}
	
--	for i = 1, nWagonsToDisable[Hydra.GameGetCurrentWorldTier()] do
--		local idWagon = T_WAGONS[i]
--		Hydra.ActorsEnableLinkedByActorInGroup(idWagon, "destroyedWagon")
--		Hydra.ActorDelete(idWagon)
--		Hydra.ActorDelete(T_VILLAGERS[idWagon])
--	end
	
	-- Very between 2-3 wagons that the player needs to defend
	for i = 1, Hydra.GetRandomInt(0, #T_WAGONS - 2) do
		local idWagon = T_WAGONS[i]
		Hydra.ActorsEnableLinkedByActorInGroup(idWagon, "destroyedWagon")
		Hydra.ActorDelete(idWagon)
		Hydra.ActorDelete(T_VILLAGERS[idWagon])
	end
	
	for i, wagonID in pairs(T_WAGONS) do -- clean up extra stuff, need this to get rid of encounter corpses that are always spawned in
		for i, linkedID in pairs(Hydra.ActorsLinkedByActorInGroup(wagonID, "destroyedWagon")) do
			Hydra.ActorDelete(linkedID)
		end
	end
	</pre></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>eQuestType</b></td><td><code>1</code></td></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>Hydra.Wait(0.1)
for i, idWagon in pairs(T_WAGONS) do
	if Hydra.ActorIsValid(idWagon) then
		Hydra.ActorSetDamageCap(idWagon, 0)
	end
end

local tYellIDs = playYell(Hydra.ActorsLinkedByActorInGroup(ID_THIS_QUEST, "Marker_Yell"))
Hydra.Wait(0.8)
tYellIDs = playYell(tYellIDs)
Hydra.Wait(0.5)
tYellIDs = playYell(tYellIDs)

playVillagerAnimAndIdle(S_KEY_HIDE)
Hydra.ActorPlaySimpleConversation(CONV_BATTLEFEAR, getRandomAliveVillager())

Hydra.ActorsEnableLinkedByActorInGroup(ID_THIS_QUEST, "Spawner_Holdout")
</pre></td></tr><tr><td><b>szOnExitScript</b></td><td><pre>--Disable further spawns in the event.
for i, idSpawner in pairs(T_SPAWNERS_ALL) do
	if Hydra.ActorIsValid(idSpawner) then
		Hydra.ActorDelete(idSpawner)
	end
end</pre></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>1416878047</code></td></tr></table>

</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>75</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>17</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>119</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>--if mastery was passed, have npcs call out warning and spawn a super elite
--Hydra.Wait(.05)
--if Hydra.QuestGetBonusPhaseResult() then
--	if isWagonsAlive() then
--		Hydra.ActorTriggerSpawnerUpdateLinkedToActor(ID_THIS_QUEST, "Spawner_EliteBoss")
--	end
--end

--make wagons invulnerable as soon as timer hits 0
for i, idWagon in pairs(T_WAGONS) do
	if Hydra.ActorIsAlive(idWagon) then
		Hydra.ActorSetInvulnerable(idWagon, true)
	end
end


Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwUID</b></td><td><code>18</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1056035579</code></td></tr><tr><td><b>eVariableType</b></td><td><code>8</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamGameTime</th></tr><tr><td><b>eParamType</b></td><td><code>12</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eFilterInequality</b></td><td><code>0</code></td></tr><tr><td><b>flTime</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>2308255985</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1056035637</code></td></tr><tr><td><b>eEventType</b></td><td><code>57</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(flTimeRemaining)
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>256</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>256</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)

	if Hydra.ActorIsAlive(T_VILLAGERS[idActorKilled]) then
		Hydra.ActorSetPowerEffectSize(T_VILLAGERS[idActorKilled], 25)
		Hydra.ActorPlayEffectGroup(T_VILLAGERS[idActorKilled], T_NPC_DEATH_FX[Hydra.GetRandomInt(1, #T_NPC_DEATH_FX)])
		Hydra.ActorKill(T_VILLAGERS[idActorKilled], DEATH_CRUSHED)
	end


end
</pre></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>67</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1536169470</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>3</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td>