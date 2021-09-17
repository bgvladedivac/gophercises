package main


import (
	"fmt"
	"log"
	"strconv"
	"os"
	"github.com/boltdb/bolt"
)


type TaskManager struct {
	PI *PersistanceInteractor
}

func CreateNewTaskManager(pi *PersistanceInteractor) *TaskManager {
	return &TaskManager{pi}
}

func (t *TaskManager) addTask(db *bolt.DB, taskName string) {
	t.PI.Update(db, taskName)
}

func (t *TaskManager) listTasks(db *bolt.DB) {
	t.PI.GetAll(db)
}

func (t *TaskManager) doTask(db *bolt.DB, taskId string) {
	t.PI.DoTask(db, taskId)
}

type PersistanceInteractor struct {
	bucketName string
}

func (pi *PersistanceInteractor) Update(db *bolt.DB, taskName string) {
	dbError := db.Update(func(tx *bolt.Tx) error {
		bucket, createBucketError := tx.CreateBucketIfNotExists([]byte(pi.bucketName))
		if createBucketError != nil {
			log.Fatal("Creating of the bucket did not happen.", createBucketError)
		}

		updateId, nextSeqError := bucket.NextSequence()
		if nextSeqError != nil {
			log.Fatal("Problem with fetching the Next Sequence ID", nextSeqError)
		}

		putError := bucket.Put([]byte(strconv.Itoa(int(updateId))), []byte(taskName))
		if putError != nil {
			log.Fatal("Problem with putting the Key/Value Pair into the bucket", putError)
		}
		return nil
	})

	if dbError != nil {
		log.Fatal("Problem with updating the bucket", dbError)
	}
}

func (pi *PersistanceInteractor) GetAll(db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(pi.bucketName))
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("Task Index=%s, Task Value=%s\n", k, v)
			return nil
		})
		return nil
	})

	if err != nil {
		log.Fatal("Problem with retriving the tasks for listing", err)
	}
}

func (pi *PersistanceInteractor) DoTask(db *bolt.DB, taskId string) {
	doTaskError := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(pi.bucketName))

		doesTaskIndexExists := false
		bucket.ForEach(func(k, v []byte) error {
			if string(k) == taskId {
				doesTaskIndexExists = true
				fmt.Println("The provided Task Id exists and it would be deleted.")
			}
			return nil
		})

		if doesTaskIndexExists {
			deletionError := bucket.Delete([]byte(taskId))
			if deletionError != nil {
				log.Fatal("Bucket was created from a read-only transaction", deletionError)
			}
		} else {
			fmt.Println("The provided Task ID does not exists")
		}
		return nil
	})

	if doTaskError != nil {
		log.Fatal("Problem with Doing the task.", doTaskError)
	}
}

func main() {
	args := os.Args[1:]
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	bucketName := "tasks"
	taskManager := CreateNewTaskManager(&PersistanceInteractor{bucketName: bucketName})
	switch action := args[0]; action {
	case "add":
		task := args[1]
		taskManager.addTask(db, task)
	case "do":
		fmt.Println("This is the argument", args[1])
		taskManager.doTask(db, args[1])
		break
	case "list":
		taskManager.listTasks(db)
	default:
		fmt.Println("Activity not supported.")
	}

}

