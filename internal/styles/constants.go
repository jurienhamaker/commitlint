package styles

import (
	"github.com/jurienhamaker/commitlint/validation"
)

var ValidationStateEmoji = map[validation.ValidationState]string{
	validation.ValidationStateError:   "⛔",
	validation.ValidationStateWarning: "⚠️",
	validation.ValidationStateSuccess: "✅",
}

var ValidationStateUnicode = map[validation.ValidationState]string{
	validation.ValidationStateError:   "✖",
	validation.ValidationStateWarning: "⚠",
	validation.ValidationStateSuccess: "✓",
}

var ValidationStateStyle = map[validation.ValidationState]func(...string) string{
	validation.ValidationStateError:   ErrorTextStyle,
	validation.ValidationStateWarning: WarningTextStyle,
	validation.ValidationStateSuccess: SuccessTextStyle,
}
