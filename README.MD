# Tool to calculate natural day set

## Today() time.Time

Returns today's time to midnight

## Tomorrow(tm time.Time) time.Time

Returns tomorrow's time by adding one day to the given time

## Yesterday(tm time.Time) time.Time

Returns tomorrow's time by subtracting one day from the given time

## Between(t1, t2 time.Time) (days int)

Calculates the difference between two times in days

## Add(tm time.Time, n int) time.Time

Adds "n" days to a given time

## Sub(tm time.Time, n int) time.Time

Subtracts "n" days to a given time