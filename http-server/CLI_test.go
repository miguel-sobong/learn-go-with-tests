package poker_test

import (
	"bytes"
	"fmt"
	"io"
	poker "learn-go-with-tests/http-server"
	"strings"
	"testing"
	"time"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := userSends("7", "Chris wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameWinner(t, game.FinishedWith, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := userSends("7", "Cleo wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameWinner(t, game.FinishedWith, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		in := userSends("7", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt
		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		AssertGameStartedWith(t, game, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userSends("Pies")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadNumberOfPlayerInputErrMsg)
		assertGameShouldNotStart(t, game)
	})

	t.Run("it prints an error when the accepted format is not met and does not start the game", func(t *testing.T) {
		in := userSends("7", "Chris is the winner")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerWinFormatInputErrMsg)
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameWinner(t testing.TB, got, winner string) {
	t.Helper()

	if got != winner {
		t.Errorf("Expected %q to win, got %q instead", winner, got)
	}
}

func userSends(inputs ...string) *strings.Reader {
	messages := strings.Join(inputs, "\n")
	return strings.NewReader(messages)
}

func assertGameShouldNotStart(t testing.TB, game *poker.GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func AssertGameStartedWith(t testing.TB, game *poker.GameSpy, expected int) {
	if game.StartedWith != expected {
		t.Errorf("wanted Start called with %d but got %d", expected, game.StartedWith)
	}
}