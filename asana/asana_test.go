package asana

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type AsanaArgs struct {
	Name      string `json:"name,omitempty"`
	Workspace string `json:"workspace,omitempty"`
	Public    bool   `json:"public,omitempty"`
	Assignee  string `json:"assignee,omitempty"`
	TaskID    string `json:"taskId,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
	Project   string `json:"id,omitempty"`
}

type NamedAndIDdEntity struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

var (
	accessToken = os.Getenv("ASANA_ACCESS_TOKEN")
	workspaceID = os.Getenv("ASANA_WORKSPACE_ID")
	projectID   = os.Getenv("ASANA_PROJECT_ID")
	taskID      = os.Getenv("ASANA_TASK_ID")
	assignee    = os.Getenv("ASANA_ASSIGNEE")
)

var _ = Describe("Create project in Asana without access token", func() {

	asana := AsanaArgs{Name: "Test Project", Workspace: workspaceID, Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana without workspace ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Name: "Test Project", Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Name: "Test Project", Workspace: workspaceID, Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create task in project Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{Name: "Test Task", Workspace: workspaceID, Assignee: assignee}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create task in project Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create task in project Asana without workspace ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Name: "Test Task", Assignee: assignee}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create task in project Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Name: "Test", Workspace: workspaceID, Assignee: assignee}
	asana.ProjectID = projectID
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(recorder, request)

	Describe("Create task in Asana", func() {
		Context("create task", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete project in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProject)
	handler.ServeHTTP(recorder, request)

	Describe("Delete project in Asana", func() {
		Context("delete project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProject)
	handler.ServeHTTP(recorder, request)

	Describe("Delete project in Asana", func() {
		Context("delete project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deleteproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProject)
	handler.ServeHTTP(recorder, request)

	Describe("Delete project in Asana", func() {
		Context("delete project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete task in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{TaskID: taskID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deletetask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)
	handler.ServeHTTP(recorder, request)

	Describe("Delete task in Asana", func() {
		Context("delete task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete task in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deletetask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)
	handler.ServeHTTP(recorder, request)

	Describe("Delete task in Asana", func() {
		Context("delete task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete task in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{TaskID: taskID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/deletetask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)
	handler.ServeHTTP(recorder, request)

	Describe("Delete task in Asana", func() {
		Context("delete task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List task in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTask)
	handler.ServeHTTP(recorder, request)

	Describe("List task in Asana", func() {
		Context("list task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List task in Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTask)
	handler.ServeHTTP(recorder, request)

	Describe("List task in Asana", func() {
		Context("list task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List task in Asana with invalid Workspace Id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Workspace: workspaceID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTask)
	handler.ServeHTTP(recorder, request)

	Describe("List task in Asana", func() {
		Context("list task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List task in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Workspace: workspaceID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTask)
	handler.ServeHTTP(recorder, request)

	Describe("List task in Asana", func() {
		Context("list task", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List workspaces in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listworkspace", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListWorkspace)
	handler.ServeHTTP(recorder, request)

	Describe("List workspaces in Asana", func() {
		Context("list workspaces", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List workspaces in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listworkspace", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListWorkspace)
	handler.ServeHTTP(recorder, request)

	Describe("List workspaces in Asana", func() {
		Context("list workspaces", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find task from Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{TaskID: taskID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindTask)
	handler.ServeHTTP(recorder, request)

	Describe("Find task in Asana", func() {
		Context("find task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find task from Asana with invaid param", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindTask)
	handler.ServeHTTP(recorder, request)

	Describe("Find task in Asana", func() {
		Context("find task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find task from Asana without Task ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{TaskID: ""}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindTask)
	handler.ServeHTTP(recorder, request)

	Describe("Find task in Asana", func() {
		Context("find task", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find task from Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{TaskID: taskID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindTask)
	handler.ServeHTTP(recorder, request)

	Describe("Find task in Asana", func() {
		Context("find task", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find project from Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindProject)
	handler.ServeHTTP(recorder, request)

	Describe("Find project in Asana", func() {
		Context("find project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find project from Asana with invaid param", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindProject)
	handler.ServeHTTP(recorder, request)

	Describe("Find project in Asana", func() {
		Context("find project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find project from Asana without project ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{ProjectID: ""}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindProject)
	handler.ServeHTTP(recorder, request)

	Describe("Find project in Asana", func() {
		Context("find project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find project from Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindProject)
	handler.ServeHTTP(recorder, request)

	Describe("Find project in Asana", func() {
		Context("find project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update project in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{Project: projectID, Name: "New Updated Project"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Update project in Asana", func() {
		Context("update project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update project in Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Update project in Asana", func() {
		Context("update project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update project in Asana without project ID ", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Name: "New Updated Project"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Update project in Asana", func() {
		Context("update project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{Project: projectID, Name: "New Updated Project"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Update project in Asana", func() {
		Context("update project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List tasks from project in Asana without access token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listprojecttasks", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProjectTasks)
	handler.ServeHTTP(recorder, request)

	Describe("List tasks from project in Asana", func() {
		Context("list tasks", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List tasks from project in Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listprojecttasks", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProjectTasks)
	handler.ServeHTTP(recorder, request)

	Describe("List tasks from project in Asana", func() {
		Context("list tasks", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List tasks from project in Asana with invalid project Id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listprojecttasks", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProjectTasks)
	handler.ServeHTTP(recorder, request)

	Describe("List tasks from project in Asana", func() {
		Context("list tasks", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List tasks from project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	asana := AsanaArgs{ProjectID: projectID}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listprojecttasks", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProjectTasks)
	handler.ServeHTTP(recorder, request)

	Describe("List tasks from project in Asana", func() {
		Context("list tasks", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Subscribe asana for task without token", func() {

	os.Setenv("ACCESS_TOKEN", "")

	data := DataArgs{WorkspaceID: workspaceID, Existing: true}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeTasks)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Subscribe asana for task without any id's", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	data := DataArgs{WorkspaceID: "", ProjectID: "", Existing: true}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeTasks)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Subscribe asana for task", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	data := DataArgs{WorkspaceID: workspaceID, Existing: true}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeTasks)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
