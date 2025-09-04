# Task Tracker CLI

Solution for roadmap.sh challenge

Project URL: https://roadmap.sh/projects/task-tracker?fl=1

# Adding a new task
go run cmd/main.go add "Buy groceries"
# Updating and deleting tasks
go run cmd/main.go update 1 "Buy groceries and cook dinner"

go run cmd/main.go delete 1
# Marking a task as in progress or done
go run cmd/main.go mark-in-progress 1

go run cmd/main.go mark-done 1
# Listing all tasks
go run cmd/main.go list
# Listing tasks by status
go run cmd/main.go list done

go run cmd/main.go list todo

go run cmd/main.go list in-progress
