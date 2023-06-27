<h1>data/base/meta/Quest/EH_WoundedDruid.qst</h1><table><tr><th colspan="100%">Metadata</th></tr><tr><td><b>Name</b></td><td>data/base/meta/Quest/EH_WoundedDruid.qst</td></tr><tr><td><b>Type</b></td><td>QuestDefinition</td></tr><tr><td><b>SNO ID</b></td><td>1293359</td></tr></table>

<table><tr><th colspan="100%">Fields</th></tr><tr><td><b>unk_43f3849</b></td><td><code>0</code></td></tr><tr><td><b>unk_f956a05</b></td><td><code>1</code></td></tr><tr><td><b>unk_8881b0e</b></td><td><code>20</code></td></tr><tr><td><b>ePlayerQuestType</b></td><td><code>-1</code></td></tr><tr><td><b>eRepeatType</b></td><td><code>0</code></td></tr><tr><td><b>unk_48a2b16</b></td><td><code>-1</code></td></tr><tr><td><b>vecStartLocation</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>arFollowers</b></td><td></td></tr><tr><td><b>arQuestDungeons</b></td><td></td></tr><tr><td><b>unk_af3a4c1</b></td><td></td></tr><tr><td><b>eQuestType</b></td><td><code>1</code></td></tr><tr><td><b>arReputationRewards</b></td><td></td></tr><tr><td><b>unk_ff5c704</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_b89b77f</b></td><td><code>0</code></td></tr><tr><td><b>eBountyType</b></td><td><code>-1</code></td></tr><tr><td><b>unk_46e3956</b></td><td></td></tr><tr><td><b>szOnAbandonScript</b></td><td><code></code></td></tr><tr><td><b>eVignetteType</b></td><td><code>0</code></td></tr><tr><td><b>unk_14dee1b</b></td><td><code>0</code></td></tr><tr><td><b>unk_c18cabd</b></td><td><code>0</code></td></tr><tr><td><b>eBountyTier</b></td><td><code>-1</code></td></tr><tr><td><b>unk_942bcdb</b></td><td><code>1</code></td></tr><tr><td><b>unk_b43b442</b></td><td></td></tr><tr><td><b>flParticipationRadius</b></td><td><code>34</code></td></tr><tr><td><b>unk_b83e7b1</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>eEventQuestType</b></td><td><code>1</code></td></tr><tr><td><b>unk_c2e8448</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_d2181f0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwNextUID</b></td><td><code>166</code></td></tr><tr><td><b>unk_79f6e17</b></td><td><code>0</code></td></tr><tr><td><b>unk_313dbf6</b></td><td><code>0</code></td></tr><tr><td><b>unk_834fdbf</b></td><td><code>4294967295</code></td></tr><tr><td><b>arRequiredReputations</b></td><td></td></tr><tr><td><b>unk_d060a69</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>szOnEventDespawnScript</b></td><td><code></code></td></tr><tr><td><b>unk_2aa5f20</b></td><td></td></tr><tr><td><b>szOnStartupScript</b></td><td><pre>ID_THIS_QUEST = Hydra.QuestGetOwner()

ID_DRUID_PROXY = Hydra.ActorSpawnActor(ID_PROXY, ID_DRUID, true, 0, 0)
Hydra.BuffApplyToTarget(POWER_DRUID_BUFF, "HOLD_AREA", ID_DRUID_PROXY)
Hydra.ActorSetHealthBarVisibility(ID_DRUID, Hydra.HPBAR_VISIBILITY_FORCE_HIDE)

nAnimalsReturned = 0
druidAlive = true</pre></td></tr><tr><td><b>arQuestPhases</b></td><td><table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwUID</b></td><td><code>6</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)
end
</pre></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>2598167221</code></td></tr><tr><td><b>szString</b></td><td><code>PlayerInLevelArea</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamSelf</th></tr><tr><td><b>dwType</b></td><td><code>1210992754</code></td></tr><tr><td><b>eParamType</b></td><td><code>11</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>50</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>49</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr><tr><td><b>unk_d17aff0</b></td><td><code>76</code></td></tr></table>


</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamPlayer</th></tr><tr><td><b>dwType</b></td><td><code>106673333</code></td></tr><tr><td><b>eParamType</b></td><td><code>7</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>bIsPlayer</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1584348101</code></td></tr><tr><td><b>eVariableType</b></td><td><code>6</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>143506636</code></td></tr><tr><td><b>eEventType</b></td><td><code>44</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>48</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorPlayer, sLevelAreaName)
	if sLevelAreaName == S_LEVEL_AREA then
		Hydra.EventSendToActor("PlayerInLevelArea", idSelf, idSelf)
	end
end
</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	if #Hydra.PlayerActorsInLevelArea(S_LEVEL_AREA) > 0 then
		Hydra.EventSendToActor("PlayerInLevelArea", idSelf, idSelf)
	end
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>56</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>7</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwFlags</b></td><td><code>2</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwUID</b></td><td><code>11</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>123</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>2106471111</code></td></tr></table>

</td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>szGroup</b></td><td><code>MonsterTrigger</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>2280539036</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
	if (not Hydra.ActorIsInPrefabScopedGroup(idActor, "EH_WoundedDruid_Monsters")) then
		Hydra.ActorAddToPrefabScopedGroup(idActor, "EH_WoundedDruid_Monsters")
	end
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>124</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>125</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>126</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>47</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>2106471111</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>szGroup</b></td><td><code>MonsterTrigger</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>2280539036</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
	if (Hydra.ActorIsInPrefabScopedGroup(idActor, "EH_WoundedDruid_Monsters")) then
		Hydra.ActorRemoveFromPrefabScopedGroup(idActor, "EH_WoundedDruid_Monsters")
	end
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>2</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>2</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>2</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	local sChest = ACTOR_CHEST_REWARD

	if Hydra.QuestGetBonusPhaseResult() then
		sChest = ACTOR_CHEST_REWARD_HEROIC
	end
	
	Events.SpawnRewardChest(CHEST_LOCATION, sChest)
end
</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>26</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>25</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>27</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwFlags</b></td><td><code>4</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>27</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>157</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>DeerTrigger</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3452001204</code></td></tr></table>


</td></tr></table>


</td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>3</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>156</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>158</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwUID</b></td><td><code>159</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(sEventName, idActor, nUserParam)
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>64</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>1212553832</code></td></tr><tr><td><b>eEventType</b></td><td><code>30</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>334654326</code></td></tr><tr><td><b>eVariableType</b></td><td><code>7</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamString</th></tr><tr><td><b>dwType</b></td><td><code>234160671</code></td></tr><tr><td><b>eParamType</b></td><td><code>6</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwStringHash</b></td><td><code>825911883</code></td></tr><tr><td><b>szString</b></td><td><code>BONUS_FAILED</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2344914334</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>43</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>84</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>eType</b></td><td><code>4</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>dwUID</b></td><td><code>76</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>74</code></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>75</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2729414872</code></td></tr><tr><td><b>eVariableType</b></td><td><code>23</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>782324564</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2537675485</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>2414055017</code></td></tr><tr><td><b>eEventType</b></td><td><code>42</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(szScopedGroup, PrefabID)
end
</pre></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>79</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>77</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(flGameTime)
	Hydra.ActorPlaySimpleConversation(CONV_CALLOUT, ID_DRUID)
end</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3510011121</code></td></tr><tr><td><b>eEventType</b></td><td><code>12</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamTimeInterval</th></tr><tr><td><b>dwType</b></td><td><code>3550172668</code></td></tr><tr><td><b>eParamType</b></td><td><code>14</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eFilterInequality</b></td><td><code>0</code></td></tr><tr><td><b>flTimeInterval</b></td><td><code>20</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>3510011063</code></td></tr><tr><td><b>eVariableType</b></td><td><code>10</code></td></tr></table>


</td></tr></table>

</td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>78</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>84</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorNPC, sConversation, idActorPlayer)
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>81</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2304262567</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>122819896</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>Druid</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1739417742</code></td></tr><tr><td><b>eVariableType</b></td><td><code>4</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamConversation</th></tr><tr><td><b>eParamType</b></td><td><code>10</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>snoConversation</b></td><td><a href="#UKNOWN">[DT_SNO] Conversation: "LE_WoundedDruid_Start"</a></td></tr><tr><td><b>dwType</b></td><td><code>2087454851</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1387993843</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>4035304214</code></td></tr><tr><td><b>eEventType</b></td><td><code>9</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>80</code></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code>Hydra.QuestAdvancePhase()</code></td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwUID</b></td><td><code>79</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code>Hydra.ActorSetTeam(ID_DRUID, Hydra.TEAM_SCENERY)</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>local deerProxy = Hydra.ActorSpawnActor(ID_PROXY, ACTIVE_SPIRIT, true, 0, 0)
Hydra.ActorFadeAndDelete(ACTIVE_SPIRIT, 2)
Hydra.Wait(0.7)
	
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.BuffApplyToTarget(POWER_SPIRIT_DEER, "RESCUE", deerProxy)
Hydra.Wait(1.3)

-- Trigger champions here
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_Champion_Spawner")
Hydra.Wait(1)


Hydra.ActorDelete(deerProxy)
Hydra.QuestAdvancePhase()</pre></td></tr><tr><td><b>dwUID</b></td><td><code>96</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>132</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>97</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>1</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3452001204</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>DeerTrigger</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
end</pre></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>dwFlags</b></td><td><code>66</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>97</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>102</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr></table>

</td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>101</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>137</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>148</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>138</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>8</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamGameTime</th></tr><tr><td><b>dwType</b></td><td><code>2308255985</code></td></tr><tr><td><b>eParamType</b></td><td><code>12</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eFilterInequality</b></td><td><code>0</code></td></tr><tr><td><b>flTime</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1056035579</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1056035637</code></td></tr><tr><td><b>eEventType</b></td><td><code>57</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(flTimeRemaining)
end
</pre></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


</td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>90</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>-- Select random spirit animals to be activated
if (RANDOM_SETUP == 1 or RANDOM_SETUP == 2) then
	ACTIVE_SPIRIT = SPIRIT_TURTLES_LEFT[1]
	InitializeSpiritAnimal(SPIRIT_TURTLES_LEFT, 1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_RIGHT, -1)
elseif (RANDOM_SETUP == 3 or RANDOM_SETUP == 4) then
	ACTIVE_SPIRIT = SPIRIT_TURTLES_MIDDLE[1]
	InitializeSpiritAnimal(SPIRIT_TURTLES_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_MIDDLE, 1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_RIGHT, -1)
else
	ACTIVE_SPIRIT = SPIRIT_TURTLES_RIGHT[1]
	InitializeSpiritAnimal(SPIRIT_TURTLES_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_TURTLES_RIGHT, 1)
end

Hydra.ActorAddToPrefabScopedGroup(ACTIVE_SPIRIT, "LE_WoundedDruid_Animals")
InitializeSpiritAnimal(SPIRIT_AMBIENT, -1)

-- Spawn monsters from random spawners
local spawners = Hydra.ActorsInPrefabScopedGroup("EH_Monster_Spawner")
if (#spawners > 0) then
	local spawnersToActivate = 7 -- out of the ~20 spawners in total
	local spawnersActivated = 0
	for i = 1, #spawners, 1 do
		local randomBool = Hydra.GetRandomInt(1, #spawners) <= spawnersToActivate
		if ((i > #spawners - spawnersToActivate + spawnersActivated or randomBool) and spawnersActivated < spawnersToActivate) then
			spawnersActivated = spawnersActivated + 1	
			Hydra.ActorTriggerSpawnerUpdate(spawners[i])
		end
	end
end

-- move temp spirit to final spot
Hydra.ActorEnable(SPIRIT_TURTLES_DIRECTOR)
Hydra.QActorMoveToLocationConstantSpeed(SPIRIT_TURTLES_DIRECTOR, ACTIVE_SPIRIT, 1, false)
Hydra.Wait(2)
Hydra.ActorFadeAndDelete(SPIRIT_TURTLES_DIRECTOR, 2)</pre></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>84</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>1</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>dwFlags</b></td><td><code>16</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>1</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>
Hydra.Wait(0.5)

if Hydra.QuestGetBonusPhaseResult() then
	-- spawn the boss if we completed the mastery
	Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_SuperElite_Spawner")
	
	local superElite = Hydra.ActorsInPrefabScopedGroup("EH_Elites")[1]
	Hydra.QuestSetCustomParamFromActor(1, superElite)

end</pre></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>117</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>27</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>118</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>42</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2729414872</code></td></tr><tr><td><b>eVariableType</b></td><td><code>23</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1948944658</code></td></tr></table>

</td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2537675485</code></td></tr><tr><td><b>eVariableType</b></td><td><code>11</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>2414055017</code></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(szScopedGroup, PrefabID)
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>116</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>60</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>nTimerDuration</b></td><td><code>30</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>132</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>1</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>local deerProxy = Hydra.ActorSpawnActor(ID_PROXY, ACTIVE_SPIRIT, true, 0, 0)
Hydra.ActorFadeAndDelete(ACTIVE_SPIRIT, 2)
Hydra.Wait(0.4)

Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.Wait(0.2)
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.Wait(0.1)
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_Elite_Spawner")	
Hydra.BuffApplyToTarget(POWER_SPIRIT_DEER, "RESCUE", deerProxy)
Hydra.Wait(1.3)

Hydra.ActorDelete(deerProxy)
Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>dwUID</b></td><td><code>133</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>139</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>134</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3452001204</code></td></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>DeerTrigger</code></td></tr></table>


</td></tr></table>


</td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
end</pre></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>1</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>134</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>66</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>dwUID</b></td><td><code>136</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr></table>

</td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>dwUID</b></td><td><code>135</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>148</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>8</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamGameTime</th></tr><tr><td><b>dwType</b></td><td><code>2308255985</code></td></tr><tr><td><b>eParamType</b></td><td><code>12</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eFilterInequality</b></td><td><code>0</code></td></tr><tr><td><b>flTime</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1056035579</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1056035637</code></td></tr><tr><td><b>eEventType</b></td><td><code>57</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>145</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(flTimeRemaining)
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>144</code></td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>1</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>-- Select random spirit animals to be activated
local randomID = Hydra.GetRandomInt(1,2)
if (RANDOM_SETUP == 1 or RANDOM_SETUP == 6) then
	ACTIVE_SPIRIT = SPIRIT_FOXES_MIDDLE[randomID]
	InitializeSpiritAnimal(SPIRIT_FOXES_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_FOXES_MIDDLE, randomID)
	InitializeSpiritAnimal(SPIRIT_FOXES_RIGHT, -1)
elseif (RANDOM_SETUP == 2 or RANDOM_SETUP == 4) then
	ACTIVE_SPIRIT = SPIRIT_FOXES_RIGHT[randomID]
	InitializeSpiritAnimal(SPIRIT_FOXES_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_FOXES_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_FOXES_RIGHT, randomID)
else
	ACTIVE_SPIRIT = SPIRIT_FOXES_LEFT[randomID]
	InitializeSpiritAnimal(SPIRIT_FOXES_LEFT, randomID)
	InitializeSpiritAnimal(SPIRIT_FOXES_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_FOXES_RIGHT, -1)
end

Hydra.ActorAddToPrefabScopedGroup(ACTIVE_SPIRIT, "LE_WoundedDruid_Animals")

-- move temp spirit to final spot
Hydra.ActorEnable(SPIRIT_FOXES_DIRECTOR)
Hydra.QActorMoveToLocationConstantSpeed(SPIRIT_FOXES_DIRECTOR, ACTIVE_SPIRIT, 2, false)
Hydra.Wait(3)
Hydra.ActorFadeAndDelete(SPIRIT_FOXES_DIRECTOR, 3)</pre></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>dwFlags</b></td><td><code>16</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>1</code></td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr><tr><td><b>unk_189b89b</b></td><td><code>1</code></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>1</code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>1</code></td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>local deerProxy = Hydra.ActorSpawnActor(ID_PROXY, ACTIVE_SPIRIT, true, 0, 0)
Hydra.ActorFadeAndDelete(ACTIVE_SPIRIT, 2)
Hydra.Wait(0.7)

Hydra.BuffApplyToTarget(POWER_SPIRIT_DEER, "RESCUE", deerProxy)

-- druid shouts & starts to get up
Hydra.ActorPlaySimpleConversation(CONV_BATTLE, ID_DRUID)
Hydra.ActorClearAnimSetOverride(ID_DRUID)
Hydra.ActorSetAIState(ID_DRUID, "Combat")
Hydra.ActorSetTeam(ID_DRUID, Hydra.TEAM_PLAYER_DEFAULT)
Hydra.Wait(0.5)

-- druid does AOE
Hydra.BuffRemove(POWER_DRUID_BUFF, "HOLD_AREA", ID_DRUID_PROXY)
Hydra.ActorSetHealthBarVisibility(ID_DRUID, Hydra.HPBAR_VISIBILITY_AUTOMATIC)
Hydra.ActorDelete(ID_DRUID_PROXY)
Hydra.ActorTriggerPowerAtFeet(ID_DRUID, POWER_DRUID_AOE)
Hydra.Wait(0.5)

-- boss spawns, kill all monsters outside of center	
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.Wait(0.3)
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.Wait(0.2)
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_MonsterWave_Spawner")
Hydra.Wait(0.1)
Hydra.ActorTriggerSpawnerUpdateInActorGroup("EH_SuperElite_Spawner")
Hydra.Wait(1)

local superElite = Hydra.ActorsInPrefabScopedGroup("EH_Elites")[1]
Hydra.QuestSetCustomParamFromActor(1, superElite)
Hydra.MinimapAddPing(superElite, 5)

-- advance
Hydra.QuestAdvancePhase()</pre></td></tr><tr><td><b>dwUID</b></td><td><code>140</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>116</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActor, idActorDetector)
end
</pre></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>2197444774</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>3941013568</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f1433225</th></tr><tr><td><b>idValue</b></td><td><code>0</code></td></tr><tr><td><b>szGroup</b></td><td><code>DeerTrigger</code></td></tr><tr><td><b>dwType</b></td><td><code>4047712805</code></td></tr><tr><td><b>eParamType</b></td><td><code>20</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>dwGroupHash</b></td><td><code>3452001204</code></td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>3442249050</code></td></tr><tr><td><b>eEventType</b></td><td><code>46</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>1</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>141</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><code></code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>142</code></td></tr><tr><td><b>ptLink</b></td><td></td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>143</code></td></tr><tr><td><b>dwFlags</b></td><td><code>66</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>1776065304</code></td></tr><tr><td><b>eEventType</b></td><td><code>0</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140265</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamAllowAny</th></tr><tr><td><b>dwType</b></td><td><code>1295852463</code></td></tr><tr><td><b>eParamType</b></td><td><code>0</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr></table>


</td></tr></table>


<table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>tKey</b></td><td><code>1189140251</code></td></tr><tr><td><b>eVariableType</b></td><td><code>1</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">Type_f0e6041c</th></tr><tr><td><b>dwType</b></td><td><code>4041606172</code></td></tr><tr><td><b>eParamType</b></td><td><code>43</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_5b5276a</b></td><td><table><tr><th colspan="100%">DT_GBID</th></tr><tr><td><b>__raw__</b></td><td><code>1701206978</code></td></tr></table>

</td></tr></table>


</td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>2</code></td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>141</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(idActorKiller, idActorKilled)
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr></table>


<table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>147</code></td></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>arEventFilters</b></td><td><table><tr><th colspan="100%">ScriptMessageMapFilter</th></tr><tr><td><b>eVariableType</b></td><td><code>8</code></td></tr><tr><td><b>unk_67545b</b></td><td><table><tr><th colspan="100%">ScriptEventParamGameTime</th></tr><tr><td><b>dwType</b></td><td><code>2308255985</code></td></tr><tr><td><b>eParamType</b></td><td><code>12</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>eFilterInequality</b></td><td><code>0</code></td></tr><tr><td><b>flTime</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>tKey</b></td><td><code>1056035579</code></td></tr></table>


</td></tr><tr><td><b>unk_492729e</b></td><td></td></tr><tr><td><b>unk_9927fd3</b></td><td><code>1056035637</code></td></tr><tr><td><b>eEventType</b></td><td><code>57</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start(flTimeRemaining)
end
</pre></td></tr><tr><td><b>arSubzones</b></td><td></td></tr></table>


</td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>dwUID</b></td><td><code>146</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr><tr><td><b>unk_d17aff0</b></td><td><code>148</code></td></tr></table>


</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>Hydra.QuestAdvancePhase()
</pre></td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr></table>

</td></tr></table>


</td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwUID</b></td><td><code>139</code></td></tr><tr><td><b>dwFlags</b></td><td><code>16</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>40</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>szOnEnterScript</b></td><td><pre>-- Select random spirit animals to be activated
local randomID = Hydra.GetRandomInt(1,2)
if (RANDOM_SETUP == 1 or RANDOM_SETUP == 3) then
	ACTIVE_SPIRIT = SPIRIT_DEER_RIGHT[randomID]
	InitializeSpiritAnimal(SPIRIT_DEER_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_DEER_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_DEER_RIGHT, randomID)
elseif (RANDOM_SETUP == 2 or RANDOM_SETUP == 5) then
	ACTIVE_SPIRIT = SPIRIT_DEER_MIDDLE[randomID]
	InitializeSpiritAnimal(SPIRIT_DEER_LEFT, -1)
	InitializeSpiritAnimal(SPIRIT_DEER_MIDDLE, randomID)
	InitializeSpiritAnimal(SPIRIT_DEER_RIGHT, -1)
else
	ACTIVE_SPIRIT = SPIRIT_DEER_LEFT[randomID]
	InitializeSpiritAnimal(SPIRIT_DEER_LEFT, randomID)
	InitializeSpiritAnimal(SPIRIT_DEER_MIDDLE, -1)
	InitializeSpiritAnimal(SPIRIT_DEER_RIGHT, -1)
end

Hydra.ActorAddToPrefabScopedGroup(ACTIVE_SPIRIT, "LE_WoundedDruid_Animals")

-- move temp spirit to final spot
Hydra.ActorEnable(SPIRIT_DEER_DIRECTOR)
Hydra.QActorMoveToLocationConstantSpeed(SPIRIT_DEER_DIRECTOR, ACTIVE_SPIRIT, 3, false)
Hydra.Wait(4)
Hydra.ActorFadeAndDelete(SPIRIT_DEER_DIRECTOR, 4)</pre></td></tr></table>


<table><tr><th colspan="100%">QuestPhase</th></tr><tr><td><b>unk_5d4cfc0</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_669bcf8</b></td><td><code>4294967295</code></td></tr><tr><td><b>dwPad</b></td><td><code>0</code></td></tr><tr><td><b>unk_eff642d</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>148</code></td></tr><tr><td><b>unk_188a07a</b></td><td></td></tr><tr><td><b>dwType</b></td><td><code>1662164195</code></td></tr><tr><td><b>unk_47f8481</b></td><td>Vector(0.000000, 0.000000, 0.000000)</td></tr><tr><td><b>unk_f6ded77</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_d9a8a05</b></td><td><code>4294967295</code></td></tr><tr><td><b>szOnEnterScript</b></td><td><code></code></td></tr><tr><td><b>unk_4b2de13</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_2bde7b6</b></td><td><code>0</code></td></tr><tr><td><b>nTimerDuration</b></td><td><code>0</code></td></tr><tr><td><b>unk_fc27941</b></td><td><code>0</code></td></tr><tr><td><b>unk_fab6e45</b></td><td><table><tr><th colspan="100%">MarkerHandle</th></tr><tr><td><b>nID</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_6344bd7</b></td><td><code>0</code></td></tr><tr><td><b>unk_f84da79</b></td><td><table><tr><th colspan="100%">Type_6b1c5d9c</th></tr><tr><td><b>hImageHandle</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_189b89b</b></td><td><code>0</code></td></tr><tr><td><b>unk_8280b0e</b></td><td><code>20</code></td></tr><tr><td><b>unk_287ecb5</b></td><td><code>4294967295</code></td></tr><tr><td><b>unk_d3b7ed7</b></td><td></td></tr><tr><td><b>szOnExitScript</b></td><td><code></code></td></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>arCallbackSets</b></td><td><table><tr><th colspan="100%">QuestObjectiveSet</th></tr><tr><td><b>dwUID</b></td><td><code>149</code></td></tr><tr><td><b>unk_68d4eb0</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>unk_8f432f8</b></td><td></td></tr><tr><td><b>tReward</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_b3249db</b></td><td><table><tr><th colspan="100%">QuestReward</th></tr><tr><td><b>nXPTier</b></td><td><code>0</code></td></tr><tr><td><b>nGoldTier</b></td><td><code>0</code></td></tr><tr><td><b>unk_b0fd814</b></td><td><code>0</code></td></tr><tr><td><b>unk_186d5e6</b></td><td><table><tr><th colspan="100%">Type_e12242af</th></tr><tr><td><b>unk_cea351b</b></td><td><code>-1</code></td></tr></table>

</td></tr><tr><td><b>unk_f8ed991</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_8c63f3c</b></td><td><code>0</code></td></tr><tr><td><b>ptLink</b></td><td><table><tr><th colspan="100%">Type_2b920147</th></tr><tr><td><b>unk_d17aff0</b></td><td><code>116</code></td></tr><tr><td><b>eLinkType</b></td><td><code>1</code></td></tr></table>


</td></tr><tr><td><b>arCallbacks</b></td><td><table><tr><th colspan="100%">QuestCallback</th></tr><tr><td><b>tScriptMessageMap</b></td><td><table><tr><th colspan="100%">ScriptMessageMap</th></tr><tr><td><b>unk_9927fd3</b></td><td><code>58</code></td></tr><tr><td><b>eEventType</b></td><td><code>65</code></td></tr><tr><td><b>dwListenerFlags</b></td><td><code>0</code></td></tr><tr><td><b>unk_c3ba01d</b></td><td><code>0</code></td></tr><tr><td><b>flCooldownSeconds</b></td><td><code>0</code></td></tr><tr><td><b>arEventFilters</b></td><td></td></tr><tr><td><b>unk_492729e</b></td><td></td></tr></table>

</td></tr><tr><td><b>unk_9c1ea4f</b></td><td><code>0</code></td></tr><tr><td><b>szOnCallbackCompleteScript</b></td><td><pre>function Start()
	Hydra.EventSend("BONUS_FAILED", ID_THIS_QUEST)
end
</pre></td></tr><tr><td><b>fDebugDisable</b></td><td><code>0</code></td></tr><tr><td><b>unk_47705e2</b></td><td><code>0</code></td></tr><tr><td><b>eIndicatorType</b></td><td><code>3</code></td></tr><tr><td><b>unk_a71fd1a</b></td><td><table><tr><th colspan="100%">ReputationValuePair</th></tr><tr><td><b>nValue</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arLevelAreas</b></td><td></td></tr><tr><td><b>arSubzones</b></td><td></td></tr><tr><td><b>unk_b185921</b></td><td><code>0</code></td></tr><tr><td><b>dwUID</b></td><td><code>150</code></td></tr><tr><td><b>tRegionDefault</b></td><td><table><tr><th colspan="100%">QuestCallbackRegion</th></tr><tr><td><b>eType</b></td><td><code>0</code></td></tr><tr><td><b>vPolygonPoints</b></td><td></td></tr><tr><td><b>vCenterPos</b></td><td>Vector(0.000000, 0.000000)</td></tr><tr><td><b>flRadius</b></td><td><code>0</code></td></tr><tr><td><b>bValid</b></td><td><code>0</code></td></tr><tr><td><b>unk_16e2f51</b></td><td><code>0</code></td></tr></table>

</td></tr><tr><td><b>arRegionOverrides</b></td><td></td></tr><tr><td><b>unk_6a71535</b></td><td><code>0</code></td></tr><tr><td><b>arSnonameTokens</b></td><td></td></tr><tr><td><b>unk_5342cc0</b></td><td><code>0</code></td></tr><tr><td><b>dwFlags</b></td><td><code>0</code></td></tr><tr><td><b>nNeeded</b></td><td><code>1</code></td></tr><tr><td><b>dwLinesSeconds</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>unk_a845fa9</b></td><td><code>0</code></td></tr><tr><td><b>szOnSetCompleteScript</b></td><td><pre>druidAlive = false



-- Remove spirit animal
local deerProxy = Hydra.ActorSpawnActor(ID_PROXY, ACTIVE_SPIRIT, true, 0, 0)
Hydra.ActorFadeAndDelete(ACTIVE_SPIRIT, 1)
Hydra.Wait(1)
Hydra.ActorDelete(deerProxy)

-- Druid dies & remove buff
Hydra.ActorPlaySimpleConversation(CONV_DEATH, ID_DRUID)
Hydra.ActorKill(ID_DRUID)
Hydra.BuffRemove(POWER_DRUID_BUFF, "HOLD_AREA", ID_DRUID_PROXY)
Hydra.ActorDelete(ID_DRUID_PROXY)

Hydra.Wait(0.5)

-- advance
Hydra.QuestAdvancePhase()
</pre></td></tr></table>


</td></tr><tr><td><b>unk_61c2846</b></td><td><code>0</code></td></tr></table>


</td></tr><tr><td><b>szUserFunctionsScript</b></td><td><pre>RANDOM_SETUP = Hydra.GetRandomInt(1, 6)

function InitializeSpiritAnimal(listOfSpirits, idToActivate)
	for i = 1, #listOfSpirits, 1 do
		if (i == idToActivate) then
			Hydra.BuffApplyToTarget(POWER_SPIRIT_DEER, "HOLD_AREA", listOfSpirits[i])
			Hydra.ActorSetAIState(listOfSpirits[i], "Wander")
		else
			Hydra.ActorFadeAndDelete(listOfSpirits[i], 1)
		end
	end
end</pre></td></tr><tr><td><b>szOnShutdownScript</b></td><td><pre>if (druidAlive) then
	Hydra.ActorPlaySimpleConversation(CONV_END, ID_DRUID)
end

Events.Despawn(idSelf, Hydra.DESPAWN_BEHAVIOR_EVENT_HUB_EVENT)</pre></td></tr><tr><td><b>arQuestItems</b></td><td></td></tr><tr><td><b>eInstanceQuestType</b></td><td><code>0</code></td></tr></table>

<h2>Quest Details (WIP)</h2><h3>Phase Order</h3>