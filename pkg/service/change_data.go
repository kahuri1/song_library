package service

import (
	"fmt"
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) ChangeData(input *model.Input) (*model.Input, error) {
	queryGroup, paramGroup := createRequestTableGroup(input)

	querySong, paramSong := createRequestTableSong(input)

	quest1, err := s.repo.UpdateGroup(queryGroup, paramGroup, input)
	if err != nil {
		return nil, err
	}
	quest2, err := s.repo.UpdateSong(querySong, paramSong, input)
	if err != nil {
		return nil, err
	}
	fmt.Println(quest1, quest2)
	return input, nil
}

func createRequestTableGroup(input *model.Input) (string, []interface{}) {
	queryGroup := "UPDATE groups SET "

	var updateGroup string

	var paramsGroup []interface{}

	paramCountGroup := 1

	if input.Group.Name != "" {
		updateGroup = updateGroup + fmt.Sprintf("name = $%d", paramCountGroup)
		paramsGroup = append(paramsGroup, input.Group.Name)
		paramCountGroup++
		queryGroup += fmt.Sprintf("%s WHERE group_id = $%d", updateGroup, paramCountGroup)
		paramsGroup = append(paramsGroup, input.Group.ID)
		return queryGroup, paramsGroup
	}

	return queryGroup, paramsGroup
}

func createRequestTableSong(input *model.Input) (string, []interface{}) {
	querySong := "UPDATE songs SET "
	var updateSong []string
	var paramsSong []interface{}
	paramCountSong := 1
	if input.Song.Title != "" {
		updateSong = append(updateSong, fmt.Sprintf("title = $%d", paramCountSong))
		paramsSong = append(paramsSong, input.Song.Title)
		paramCountSong++
	}

	if input.Song.ReleaseDate != "" {
		updateSong = append(updateSong, fmt.Sprintf("release_date = $%d", paramCountSong))
		paramsSong = append(paramsSong, input.Song.ReleaseDate)
		paramCountSong++

	}

	if input.Song.Lyrics != "" {
		updateSong = append(updateSong, fmt.Sprintf("lyrics = $%d", paramCountSong))
		paramsSong = append(paramsSong, input.Song.Lyrics)
		paramCountSong++

	}

	if input.Song.Link != "" {
		updateSong = append(updateSong, fmt.Sprintf("link = $%d", paramCountSong))
		paramsSong = append(paramsSong, input.Song.Link)
		paramCountSong++
	}
	if input.Song.GroupID != 0 {
		updateSong = append(updateSong, fmt.Sprintf("group_id = $%d", paramCountSong))
		paramsSong = append(paramsSong, input.Song.GroupID)
		paramCountSong++
	}

	querySong += fmt.Sprintf("%s WHERE song_id = $%d", sqlJoin(updateSong), paramCountSong)
	paramsSong = append(paramsSong, input.Song.SongID)
	return querySong, paramsSong
}

func sqlJoin(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	return stringJoin(parts, ", ")
}

func stringJoin(parts []string, sep string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += sep
		}
		result += part
	}
	return result
}
