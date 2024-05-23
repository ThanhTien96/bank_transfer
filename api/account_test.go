package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	mockdb "simplebank/db/mock"
	db "simplebank/db/sqlc"
	"simplebank/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request
	server := NewServer(store)
	recoder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID);
	fmt.Println(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.route.ServeHTTP(recoder, request)
	// check response
	require.Equal(t, http.StatusOK, recoder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		ID: utils.RandomInt(1, 1000),
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unma
}