# Export 되는 이름

Go 에서는 대문자로 시작되는 이름이 export 됩니다. 예) math 패키지의 `Pi`

package를 import 할 때, 그 패키지의 export 된 이름만 참조 할 수 있습니다.

export 되지 않은 이름은 패키지의 밖에서 접근 할 수 없습니다.