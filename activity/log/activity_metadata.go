package log

var jsonMetadata = `{
  "name": "tibco-log",
  "type": "flogo:activity",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/log",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "inputs":[
    {
      "name": "message",
      "type": "string",
      "value": ""
    },
    {
      "name": "flowInfo",
      "type": "boolean",
      "value": "false"
    },
    {
      "name": "addToFlow",
      "type": "boolean",
      "value": "false"
    }
  ],
  "outputs": [
    {
      "name": "message",
      "type": "string"
    }
  ]
}`
