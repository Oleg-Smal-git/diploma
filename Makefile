build: __build_ecs __build_oop __build_render

__build_ecs:
	@go build -o ./bin/physics/ecs -tags BUILD_ECS ./main/physics

__build_oop:
	@go build -o ./bin/physics/oop -tags BUILD_OOP ./main/physics

__build_render:
	@go build -o ./bin/graphics ./main/graphics

simulate_ecs:
	@./bin/physics/ecs

simulate_oop:
	@./bin/physics/oop

render:
	@./bin/graphics