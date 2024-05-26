package model

import (
	"context"
	"io/fs"
	"os"
	"time"

	"github.com/go-shiori/warc"
	"github.com/spf13/afero"
)

type BookmarksDomain interface {
	HasEbook(b *BookmarkDTO) bool
	HasArchive(b *BookmarkDTO) bool
	HasThumbnail(b *BookmarkDTO) bool
	GetBookmark(ctx context.Context, id DBID) (*BookmarkDTO, error)
}

type TagsDomain interface {
	GetTags(ctx context.Context) ([]TagDTO, error)
	CreateTags(ctx context.Context, tags ...TagDTO) (TagDTO, error)
	UpdateTag(ctx context.Context, tag TagDTO) (TagDTO, error)
	DeleteTag(ctx context.Context, tagID DBID) error
}

type AccountsDomain interface {
	CheckToken(ctx context.Context, userJWT string) (*Account, error)
	GetAccountFromCredentials(ctx context.Context, username, password string) (*Account, error)
	CreateTokenForAccount(account *Account, expiration time.Time) (string, error)
}

type ArchiverDomain interface {
	DownloadBookmarkArchive(book BookmarkDTO) (*BookmarkDTO, error)
	GetBookmarkArchive(book *BookmarkDTO) (*warc.Archive, error)
}

type StorageDomain interface {
	Stat(name string) (fs.FileInfo, error)
	FS() afero.Fs
	FileExists(path string) bool
	DirExists(path string) bool
	WriteData(dst string, data []byte) error
	WriteFile(dst string, src *os.File) error
}
