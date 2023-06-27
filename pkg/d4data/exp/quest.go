package exp

import (
	"encoding/json"
	"github.com/windler/dotgraph/graph"
	"strconv"
)

func nodeName(t string, uid int) string {
	return t + " " + strconv.FormatInt(int64(uid), 10)
}

type QuestDefinition struct {
	QuestPhases []QuestPhase `json:"arQuestPhases"`
	NextUID     int          `json:"dwNextUID"`
}

func (qd QuestDefinition) Graph() *graph.DotGraph {
	g := graph.New("QuestDefinition")

	g.AddDirectedEdge("START", nodeName("Phase", 0), "")
	g.AddNodeGraphPatternOptions("START", graph.DotGraphOptions{
		"shape": "diamond",
	})

	for _, phase := range qd.QuestPhases {
		for _, callbackSet := range phase.CallbackSets {
			g.AddDirectedEdge(nodeName("Phase", phase.UID), nodeName("CallbackSet", callbackSet.UID), "")

			for _, link := range callbackSet.Links {
				g.AddDirectedEdge(nodeName("CallbackSet", callbackSet.UID), nodeName("Phase", link.DestinationPhaseUID), "AdvanceQuestPhase()")
			}

			for _, callback := range callbackSet.Callbacks {
				g.AddDirectedEdge(nodeName("CallbackSet", callbackSet.UID), nodeName("Callback", callback.UID), "")
			}
		}
	}

	return g
}

type QuestPhase struct {
	UID          int                 `json:"dwUID"`
	CallbackSets []QuestObjectiveSet `json:"arCallbackSets"`
}

type QuestObjectiveSet struct {
	UID       int             `json:"dwUID"`
	Links     []Link          `json:"ptLink"`
	Callbacks []QuestCallback `json:"arCallbacks"`
}

type Link struct {
	DestinationPhaseUID int `json:"unk_d17aff0"` // dwDestinationPhaseUID
}

type QuestCallback struct {
	UID int `json:"dwUID"`
}

func ParseQuestDef(s string) (QuestDefinition, error) {
	var qd QuestDefinition
	err := json.Unmarshal([]byte(s), &qd)
	return qd, err
}
