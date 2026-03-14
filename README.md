# birthday-bot
구글 캘린더에서 일정 조회 후 슬랙에 생일자 알림

프로젝트 초기화

```bash
go mod init <프로젝트 경로>
```

* `package main`은 진입점 패키지 (spring에서 `@SpringBootApplication` 있는 클래스)
* `func main()`은 `public static void main(String[] args)`

의존성 패키지 설치
`go get -u <package 경로>`

### 포인터 (*)

* `Config`는 값 자체를 복사
* `*Config`는 Config를 가리키는 포인터 (주소) -> 참조 타입처럼 동작
* `&Config`는 Config를 가리키는 주소를 가져옴

```go
// 값으로 전달 - 복사본이 넘어감
func updateValue(c Config) {
    c.Value = "new"  // 원본 안 바뀜
}

// 포인터로 전달 - 원본 주소가 넘어감
func updateValuePtr(c *Config) {
    c.Value = "new"  // 원본 바뀜
}

func main() {
    config := Config{Key: "test", Value: "old"}
    
    updateValue(config)
    fmt.Println(config.Value)  // "old" (안 바뀜)
    
    updateValuePtr(&config)    // &로 주소 전달
    fmt.Println(config.Value)  // "new" (바뀜!)
}
```

Go는 class가 없어서 생성자 주입을 아래와 같이함.

```kotlin
class ConfigRepository (
    private val db: Database
) {
    fun findByKey(Key: String): Config? {
        ...
    }
}

val repository = ConfigRepository(db)
```

```go
// struct로 필드만 정의 (class의 property)
type ConfigRepository struct {
    db *gorm.DB
}

// 팩토리 함수가 생성자 주입을 해주는 역할
func NewConfigRepository(db *gorm.DB) *ConfigRepository {
    return &ConfigRepository{db: db}
}

// 메서드는 struct에 붙이는 함수
func (r *ConfigRepository) FindByKey(key string) (*Config, error) {
    // r이 this 역할을 함
}
```

### make 함수
slice, map, channel을 초기화하는 함수

slice에서 make
```go
// 방법 1: 빈 슬라이스 (길이 0)
users := []string{}

// 방법 2: mke로 길이 지정
users := make([]string, 5) // 길이 5, 용량 5
users := make([]string, 0, 10) // 길이 0, 용량 10
```

#### make를 사용하는 이유
```go
	slackUsers := make([]*domain.SlackUser, len(dto))

	for i, item := range dto {
		slackUsers[i] = domain.NewSlackUser(...) // 인덱스로 할당
	}
```
미리 길이를 정해두면 인덱스 바로 접근 가능함
만약에 make를 사용하지 않으면
```go
slackUsers := []*domain.SalckUser{} // 빈 slice

for _, item range dto {
    slackUser := domain.NewSalckUser(...)
    slackUsers = append(...)// append로 추가해줘야 함
}
```

### 성능 차이
* `make([]T, n)` : 메모리 한 번에 할당, 인덱스 접근
* `append` : 용량 초과 시 메모리 재할당 발생

길이를 미리 알 수 있다면 `make` 가 더 효율적이다.