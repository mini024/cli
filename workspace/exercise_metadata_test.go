package workspace

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExerciseMetadata(t *testing.T) {
	dir, err := os.MkdirTemp("", "solution")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	em1 := &ExerciseMetadata{
		Track:        "a-track",
		ExerciseSlug: "bogus-exercise",
		ID:           "abc",
		URL:          "http://example.com",
		Handle:       "alice",
		IsRequester:  true,
		Dir:          dir,
	}
	err = em1.Write(dir)
	assert.NoError(t, err)

	em2, err := NewExerciseMetadata(dir)
	assert.NoError(t, err)
	assert.Nil(t, em2.SubmittedAt)
	assert.Equal(t, em1, em2)

	ts := time.Date(2000, 1, 2, 3, 4, 5, 6, time.UTC)
	em2.SubmittedAt = &ts

	err = em2.Write(dir)
	assert.NoError(t, err)

	em3, err := NewExerciseMetadata(dir)
	assert.NoError(t, err)
	assert.Equal(t, em2, em3)
}

func TestSuffix(t *testing.T) {
	testCases := []struct {
		metadata ExerciseMetadata
		suffix   string
	}{
		{
			metadata: ExerciseMetadata{
				ExerciseSlug: "bat",
				Dir:          "",
			},
			suffix: "",
		},
		{
			metadata: ExerciseMetadata{
				ExerciseSlug: "bat",
				Dir:          "/path/to/bat",
			},
			suffix: "",
		},
		{
			metadata: ExerciseMetadata{
				ExerciseSlug: "bat",
				Dir:          "/path/to/bat-2",
			},
			suffix: "2",
		},
		{
			metadata: ExerciseMetadata{
				ExerciseSlug: "bat",
				Dir:          "/path/to/bat-200",
			},
			suffix: "200",
		},
	}

	for _, tc := range testCases {
		testName := "Suffix of '" + tc.metadata.Dir + "' should be " + tc.suffix
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.suffix, tc.metadata.Suffix(), testName)
		})
	}
}

func TestExerciseMetadataString(t *testing.T) {
	testCases := []struct {
		metadata ExerciseMetadata
		desc     string
	}{
		{
			metadata: ExerciseMetadata{
				Track:        "elixir",
				ExerciseSlug: "secret-handshake",
				Handle:       "",
				Dir:          "",
			},
			desc: "elixir/secret-handshake",
		},
		{
			metadata: ExerciseMetadata{
				Track:        "cpp",
				ExerciseSlug: "clock",
				Handle:       "alice",
				IsRequester:  true,
			},
			desc: "cpp/clock",
		},
		{
			metadata: ExerciseMetadata{
				Track:        "cpp",
				ExerciseSlug: "clock",
				Handle:       "alice",
				IsRequester:  true,
				Dir:          "/path/to/clock-2",
			},
			desc: "cpp/clock (2)",
		},
		{
			metadata: ExerciseMetadata{
				Track:        "fsharp",
				ExerciseSlug: "hello-world",
				Handle:       "bob",
				IsRequester:  false,
			},
			desc: "fsharp/hello-world by @bob",
		},
		{
			metadata: ExerciseMetadata{
				Track:        "haskell",
				ExerciseSlug: "allergies",
				Handle:       "charlie",
				IsRequester:  false,
				Dir:          "/path/to/allergies-2",
			},
			desc: "haskell/allergies (2) by @charlie",
		},
	}

	for _, tc := range testCases {
		testName := "should stringify to '" + tc.desc + "'"
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.desc, tc.metadata.String())
		})
	}
}

func TestExerciseMetadataFilePermissions(t *testing.T) {
	dir, err := os.MkdirTemp("", "solution")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	em := &ExerciseMetadata{
		Track:        "a-track",
		ExerciseSlug: "bogus-exercise",
		ID:           "abc",
		URL:          "http://example.com",
		Handle:       "alice",
		IsRequester:  true,
		Dir:          dir,
	}
	
	err = em.Write(dir)
	assert.NoError(t, err)

	// Check the metadata file permissions
	metadataPath := filepath.Join(dir, ".exercism", "metadata.json")
	info, err := os.Stat(metadataPath)
	assert.NoError(t, err)
	
	// The file should have 0600 permissions (user read/write only)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm(), "Metadata file should have 0600 permissions")
	
	// Check the .exercism directory permissions
	exercismDir := filepath.Join(dir, ".exercism")
	dirInfo, err := os.Stat(exercismDir)
	assert.NoError(t, err)
	
	// The directory should have 0755 permissions
	assert.Equal(t, os.FileMode(0755), dirInfo.Mode().Perm(), ".exercism directory should have 0755 permissions")
}
