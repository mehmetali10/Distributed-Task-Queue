package userhelper

import (
	"context"
	"mid/core/auth"
	customerrors "mid/core/errors"
	"mid/core/var/common"
)

// ExctractUserFromContext extracts the user claims from the provided context.
// If the user claims are not present or incomplete, it returns an error.
func ExctractUserFromContext(ctx context.Context) (*auth.User, error) {

	claims := ctx.Value(common.LabelUser).(*auth.User)
	if claims.Name == "" {
		return &auth.User{}, &customerrors.NoContextFound{}
	}
	return claims, nil

}
