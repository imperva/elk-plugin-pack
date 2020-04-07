package models

import (
	"time"
)

type Header struct {
    Vendor            string  `json:"vendor"`
    Product           string  `json:"product"`
    Type              string  `json:"type"`
    Subtype           string  `json:"subtype"`
    Version           string  `json:"template_version"`
    Date              string  `json:"template_date"`
}

type Event struct {
    DbUser            string  `json:"db-user"`
    EventID           string  `json:"event-id"`
    EventType         string  `json:"event-type"`
    UserGroup         string  `json:"user-group"`
    UserAuthenticated string  `json:"user-authenticated"`
    ApplicationUser   string  `json:"application-user"`
    SourceApplication string  `json:"source-application"`
    OsUser            string  `json:"os-user"`
    OsUserChain       string  `json:"os-user-chain"`
    RawQuery          string  `json:"raw-query"`
    ParsedQuery       string  `json:"parsed-query"`
    BindVariables     string  `json:"bind-variables"`
    SqlError          string  `json:"sql-error"`
    AffectedRows      int     `json:"affected-rows"`
    ServiceType       string  `json:"service-type"`
    MxIp              string  `json:"mx-ip"`
    GwIp              string  `json:"gw-ip"`
    AgentName         string  `json:"agent-name"`
    DbSchemaPair      string  `json:"db-schema-pair"`
    DbSchemaName      string  `json:"db-schema-name"`
    StoredProcedure   string  `json:"stored-procedure"`
    OperationName     string  `json:"operation-name"`
    DatabaseName      string  `json:"database-name"`
    TableGroupName    string  `json:"table-group-name"`
    IsSensitive       string  `json:"is-sensitive"`
    IsPrivileged      string  `json:"is-privileged"`
    ObjectType        string  `json:"object-type"`
    ObjectsList       string  `json:"objects-list"`
    ResponseSize      int     `json:"response-size"`
    ResponseTime      int     `json:"response-time"`
    ExceptionOccured  string  `json:"exception-occurred"`
    ExceptionMessage  string  `json:"exception-message"`
    QueryGroup        string  `json:"query-group"`
}

type Message struct {
    Header            Header  `json:"header"`
    EventTime         string  `json:"event-time"`
    DayOfWeek         string  `json:"day-of-week"`
    HourOfDay         string  `json:"hour-of-day"`
    Version           string  `json:"SecureSphere-Version"`
    MxIp              string  `json:"mx-ip"`
    GwIp              string  `json:"gw-ip"`
    DestIp            string  `json:"dest-ip"`
    DestPort          string  `json:"dest-port"`
    HostName          string  `json:"host-name"`
    SourceIp          string  `json:"source-ip"`
    SourcePort        string  `json:"source-port"`
    Protocol          string  `json:"protocol"`
    Policy            string  `json:"policy"`
    ServerGroup       string  `json:"server-group"`
    ServerName        string  `json:"service-name"`
    ApplicationName   string  `json:"application-name"`
    Event             Event   `json:"event"`
    Timestamp         time.Time  `json:"@timestamp"`
}