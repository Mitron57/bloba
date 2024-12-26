package handlers

import (
    "encoding/json"
    "errors"
    "net/http"
    "seems/application/services"
    "seems/domain/dto"
    "seems/domain/interfaces"
)

type UserHandler struct {
    service interfaces.UserService
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    credentials := dto.Credentials{}
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        err, _ := json.Marshal(map[string]string{"error": err.Error()})
        w.Write(err)
        return
    }
    user, err := u.GetUserByCredentials(&credentials)
    if err != nil {
        if errors.Is() {
            w.WriteHeader(http.StatusNotFound)
            w.Write()
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, user)
}
