initialize:
	@make __build_initialize
	@mkdir -p ./buff
	@./bin/initialize

render_ecs: __build_export_ecs __build_render __ecs __render

profile_ecs: __build_profile_ecs __ecs __profile

render_oop: __build_export_oop __build_render __oop __render

profile_oop: __build_profile_oop __oop __profile

__build_initialize:
	@go build -o ./bin/init ./init

__build_export_ecs:
	@go build -o ./bin/physics -tags BUILD_ECS,BUILD_EXPORT ./main/physics

__build_profile_ecs:
	@go build -o ./bin/physics -tags BUILD_ECS,BUILD_PROFILE ./main/physics

__build_export_oop:
	@go build -o ./bin/physics -tags BUILD_ECS,BUILD_EXPORT ./main/physics

__build_profile_oop:
	@go build -o ./bin/physics -tags BUILD_ECS,BUILD_PROFILE ./main/physics

__build_render:
	@go build -o ./bin/graphics ./main/render

__ecs:
	@./bin/physics

__oop:
	@./bin/physics

__render:
	@./bin/graphics

__profile:
	@go tool pprof -png ./bin/physics ./buff/cpu-profile > ./buff/cpu-profile.png
	@go tool pprof -png ./bin/physics ./buff/memory-profile > ./buff/memory-profile.png
