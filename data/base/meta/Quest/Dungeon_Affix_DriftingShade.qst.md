<h1>data/base/meta/Quest/Dungeon_Affix_DriftingShade.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/Dungeon_Affix_DriftingShade.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>573824</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_c2e8448</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwNextUID</b></td><td><code>34</code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>unk_48a2b16</b></td><td><code>-1</code></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_8881b0e</b></td><td><code>20</code></td></tr><tr><td><b>arRequiredReputations</b></td><td></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szOnShutdownScript</b></td><td><code></code></td></tr><tr><td><b>arFollowers</b></td><td></td></tr><tr><td><b>arQuestItems</b></td><td></td></tr><tr><td><b>unk_b43b442</b></td><td></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>unk_2aa5f20</b></td><td></td></tr><tr><td><b>eEventQuestType</b></td><td><code>0</code></td></tr><tr><td><b>eRepeatType</b></td><td><code>0</code></td></tr><tr><td><b>unk_d2181f0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_14dee1b</b></td><td><code>0</code></td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>unk_b83e7b1</b></td><td><code>0</code></td></tr><tr><td><b>szUserFunctionsScript</b></td><td><pre>function UF_Spawn_Shade(idTargetPlayer)
	--get a random location that fits requirements around the player, and check LOS. Only spawn at a point that player LOS check is successful
	local ptTargetLoc = Hydra.ActorGetPosition(idTargetPlayer);
	--validate player position on navmesh, if not, like a bar is leaping in the air, then get the most closed position on navmesh
	if not Hydra.PositionIsOnValidNavmesh(ptTargetLoc) then 
		ptTargetLoc = Hydra.PositionSearchForNearbyNavmeshPosition(ptTargetLoc, 0, 0.5, true);
	end
	
	local idShade = Hydra.ACDID_INVALID;
	local ptSpawnLoc = Hydra.PositionSearchForNearbyNavmeshPosition(ptTargetLoc, flMinDisFromTarget, flMaxDisFromTarget);
	local ptFacing = Hydra.PositionGetFacingToPosition(ptSpawnLoc, ptTargetLoc);	
	
	--Spawn the shade any have it use the targeting power, which handles all warning and initialization for the shade and target.
	idShade = Hydra.ActorSpawnActor(S_SHADE_CORE_ACTOR, ptSpawnLoc, ptFacing, false, 0, 0);
	if not Hydra.ActorIsValid(idShade) then
		return;
	end

	Hydra.ActorTriggerPowerAtTarget(idShade, idTargetPlayer, SNO_POWER_TARAGETED);
end</pre></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>unk_af3a4c1</b></td><td></td></tr><tr><td><b>eInstanceQuestType</b></td><td><code>1</code></td></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>ePlayerQuestType</b></td><td><code>-1</code></td></tr><tr><td><b>eVignetteType</b></td><td><code>0</code></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>23</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>24</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	--initial delay before affix takes effects
	Hydra.Wait(FL_DUNGEONAFFIX_COOLUP);
	
	--keep looping to spawn shades
	while (true) do
		--=====STOP CONDITION: region is not dungeon====--
		--If the area somehow is NOT a dungeon, stop this loop. It must be some errors, or a misuseage since affixes are only supposed to be in a dungeon
		if not Hydra.SubzoneIsDungeon() then
			break;
		end
		--=====STOP CONDITION====--
		
		--get player actors
		--local idPlayerValid = Hydra.ACDID_INVALID;
		local tPlayer = Hydra.PlayerActors();
		local tViableTargets = {}
		
		for i, idPlayer in ipairs(tPlayer) do
			if Hydra.ActorIsValid(idPlayer)
				and Hydra.ActorIsAlive(idPlayer)
				and Hydra.ActorIsInCombat(idPlayer)
				and not Hydra.BuffIsOnActor(SNO_POWER_TARAGETED, "BUFF_TARGETED", idPlayer) then
				
				table.insert(tViableTargets, idPlayer)
			end
		end
		
		--=====STOP CONDITION: no valid player====--
		--[[ Comment this part of stop condition out, since for now we allow players to exit and re-enter ND. If this logic is active, in a single player game, player can exit then come back to stop the affix quest and cheese the ND.
		--if there's no valid player, break the loop
		if not Hydra.ActorIsValid(idPlayerValid) then
			--break;
		end]]--
		--=====STOP CONDITION====--
		
		--Hydra.Trace("combat players: {c_blue}"..#tViableTargets.."{/c} || Shades: {c_blue}"..#tTargetedPlayers.."{/c}");
		--====SPAWN LOGIC====--		
		if #tViableTargets == 1 then
			--User Function: spawn the shade here, mark the player with buff tag, and force it to be the target of the shade
			UF_Spawn_Shade(tViableTargets[1]);
			Hydra.Wait(FL_NOSHADETIME_MIN + (FL_NOSHADETIME_OFFSET * Hydra.GetRandomFloat())); --offset is 0~3s, total no shade time is 10~13s ,set in OnStartup
			
		elseif #tViableTargets > 1 then
				
			ShuffleTable(tViableTargets);
			--find the first player that is not chased by a shade (can tell by the buff), and spawn a shade for this player
			for i, idPlayer in ipairs (tViableTargets) do
				if not Hydra.BuffIsOnActor(SNO_POWER_TARAGETED, "BUFF_TARGETED", idPlayer) then
					--User Function: spawn the shade here, mark the player with buff tag, and force it to be the target of the shade
					UF_Spawn_Shade(idPlayer);
					break;
				end	
			end
			--for each shade spawn, wait for a small time, and make sure the interval is large enough so they don't spawn the same time
			Hydra.Wait(FL_DIFFTIME + FL_DIFFTIME_OFFSET * Hydra.GetRandomFloat()); --offset is 0~2s, total interval time is 1~3s, set in OnStartup
		else
			--if there are already enough shades, then enter cooldown. Also take a break, if no combat player then no cores. Cool down time scale with player number, scale has diminishing return on more players
			--if the player leaves the dungeon, tPlayers will become null, so use math.max to default to 1 if the player(s) is no longer present.
			--Hydra.Trace("Enter shade cool down, "..#tPlayer.." players exist, rest for "..FL_NOSHADETIME_MIN * GROUP_SCALAR[#tPlayer] + (FL_NOSHADETIME_OFFSET * math.min( Hydra.GetRandomFloat(), 0.33 )).." seconds");
			Hydra.Wait(FL_NOSHADETIME_MIN * GROUP_SCALAR[math.max(1,#tPlayer)] + (FL_NOSHADETIME_OFFSET * Hydra.GetRandomFloat())); --offset is 0~3s, total no shade time is 10~13s, also multiplied by the group number scalar, set in OnStartup
		end
		--====SPAWN LOGIC====--
		
		--give while loop a global cooldown
		Hydra.Wait(0.5)
	end -- while loop
end
</pre></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>dwUID</b></td><td><code>12</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwFlags</b></td><td><code>6</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr></table>


</td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>szOnStartupScript</b></td><td><pre>--generic affix properties
S_DUNGEON_AFFIX_NAME = "Drifting Shade";
FL_DUNGEONAFFIX_COOLUP = 24;
FL_NOSHADETIME_MIN = 10;
FL_NOSHADETIME_OFFSET = 3;
FL_DIFFTIME = 1;
FL_DIFFTIME_OFFSET = 2;

--BALANCE KNOB: cool down scalar for 1 to 4 players, index is the player number
GROUP_SCALAR = {1, 1, 0.84, 0.75};

--the shade spawn range, how far away from the targeted player
flMinDisFromTarget = 6;
flMaxDisFromTarget = 9;
</pre></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>eQuestType</b></td><td><code>2</code></td></tr><tr><td><b>unk_c18cabd</b></td><td><code>0</code></td></tr><tr><td><b>unk_942bcdb</b></td><td><code>1</code></td></tr><tr><td><b>flParticipationRadius</b></td><td><code>12</code></td></tr><tr><td><b>eBountyTier</b></td><td><code>-1</code></td></tr></table>

<h2>Quest Details (WIP)</h2><h3>Phase Order</h3>