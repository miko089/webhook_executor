package configparser

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func cmp(a *Config, b *Config) bool {
	if len(a.Tasks) != len(b.Tasks) {
		return false
	}
	for i := range a.Tasks {
		if a.Tasks[i].ID != b.Tasks[i].ID || a.Tasks[i].Command != b.Tasks[i].Command {
			return false
		}
	}
	return true
}

// TestParseTask tests unmarshalling of yaml file
func TestParseTask(t *testing.T) {
	t.Log("Testing unmarshalling of parseFile function")
	{
		testID := 0
		t.Logf("\tTest %d:\t unmarshall single task", testID)
		{
			conf := `tasks:
  - id: test1
    command: "touch test1"
  - id: test2
    command: "touch test2"
  - id: test3
    command: "touch test3 && echo \"meow meow meow\" >> test3"`
			config, err := parseFile([]byte(conf))
			shouldBe := &Config{
				Tasks: []*Task{{ID: "test1", Command: "touch test1"},
					{ID: "test2", Command: "touch test2"},
					{ID: "test3", Command: "touch test3 && echo \"meow meow meow\" >> test3"}}}
			if err != nil {
				t.Fatalf("\t%s\tShouldn't have errored : %v.", failed, err)
			}
			if !cmp(config, shouldBe) {
				t.Fatalf("\t%s\tShould be %+v, but got %+v.", failed, shouldBe, config)
			}
			t.Logf("\t%s\tSuccessful unmarshalling", success)
		}
	}
}
