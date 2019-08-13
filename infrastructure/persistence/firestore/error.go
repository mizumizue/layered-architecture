package firestore

import "golang.org/x/xerrors"

func iteratorNextFailed(detail error) error {
	return xerrors.Errorf("itr.Next failed. error: %w", detail)
}

func convertToDataFailed(detail error) error {
	return xerrors.Errorf("convert data to struct error: %w", detail)
}

func documentNotFound(detail error) error {
	return xerrors.Errorf("document not found: %w", detail)
}

func firestoreUnknownErr(detail error) error {
	return xerrors.Errorf("unknown error: %w", detail)
}
