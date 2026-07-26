package domains

type DateTimeService struct {
	clock Clock
}

func (dt *DateTimeService) CurrentUnixSeconds() int64 {
	return dt.clock.NowUnix()
}

func NewDateTimeService(c Clock) *DateTimeService {
	return &DateTimeService{clock: c}
}
