package main

import (
	"encoding/json"
	"testing"
)

var (
	fakeResults = []byte(`
{
  "zipkin-2017-03-11" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-11",
        "creation_date" : "1489190400015",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "h1JFT-9CReehz5ZjLCjCjg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-22" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-22",
        "creation_date" : "1487721600008",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "FodUaYWMQbG_Ipo3T-RPkQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-19" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-19",
        "creation_date" : "1487462400043",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "gBVNYbzpQH6PK6cPC5JK8g",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-06" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-06",
        "creation_date" : "1488758400019",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "QGkItFUdTpSH6IspbUR7eA",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-23" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-23",
        "creation_date" : "1487808000001",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "08Tay7s8TL-sdwBjkY_QkA",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-28" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-28",
        "creation_date" : "1488240000012",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "uMboVswCTtS2CXAR9EWw2g",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-05" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-05",
        "creation_date" : "1488672000014",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "bceZTOOnT_unpaScGhfomg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-08" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-08",
        "creation_date" : "1488931199993",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "O8VoByT0QwKNls9IWbFFcg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-12" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-12",
        "creation_date" : "1489277290599",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "KQD5-PauSaKHn436xLwHzg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-02" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-02",
        "creation_date" : "1488412800000",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "IDVxTkFISD-122nJZ39ZHg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-20" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-20",
        "creation_date" : "1487548799996",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "HE_HmDeDRFS1F-TH9hMpMg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-16" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-16",
        "creation_date" : "1487203200010",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "ctDd9a5pRx-ksCxJEvaa9g",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-09" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-09",
        "creation_date" : "1489017600008",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "05DALmxTS1a_WljJkDde9w",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-03" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-03",
        "creation_date" : "1488499200015",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "WyXB2d43SDObcMZLa5NmeQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-26" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-26",
        "creation_date" : "1488067200002",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "rEKFmCroSpepp7FtKYDfUg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-24" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-24",
        "creation_date" : "1487894400002",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "zVGsxC7lRWWocNzA6A8bqQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-17" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-17",
        "creation_date" : "1487289600029",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "fKnB7IhTTvinfSP1vALYSA",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-15" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-15",
        "creation_date" : "1487116800072",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "NNWg-dNWR9mS7h9mY7zUmw",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-25" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-25",
        "creation_date" : "1487980800028",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "LgU2YvkHRnGR2yl98vhTkQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-04" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-04",
        "creation_date" : "1488585600020",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "bXGHfD9UQDS0i1U4Yrvl9A",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-10" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-10",
        "creation_date" : "1489104000005",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "4jsJx-blTkO-LaEPx0ZaDg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-27" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-27",
        "creation_date" : "1488153600004",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "5cE6-1LET2i3g_xnatkkZg",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-21" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-21",
        "creation_date" : "1487635200014",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "lXKIu-ghRsSRDhpLEC8Btw",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-01" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-01",
        "creation_date" : "1488326400031",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "s1KGReQ0T9-9Th5cSpvLmQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-13" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-13",
        "creation_date" : "1489364577450",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "TP0MIZYCQlykMciNkVhVmw",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-18" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-18",
        "creation_date" : "1487376000018",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "VqRSUiYmQLuL9Gu-mwaL5Q",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-03-07" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-03-07",
        "creation_date" : "1488844800014",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "2dff0izBT3W0Mkq43-F0nQ",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  },
  "zipkin-2017-02-14" : {
    "settings" : {
      "index" : {
        "number_of_shards" : "5",
        "blocks" : {
          "write" : "false"
        },
        "provided_name" : "zipkin-2017-02-14",
        "creation_date" : "1487104024437",
        "requests" : {
          "cache" : {
            "enable" : "true"
          }
        },
        "analysis" : {
          "filter" : {
            "traceId_filter" : {
              "type" : "pattern_capture",
              "preserve_original" : "true",
              "patterns" : [
                "([0-9a-f]{1,16})$"
              ]
            }
          },
          "analyzer" : {
            "traceId_analyzer" : {
              "filter" : "traceId_filter",
              "type" : "custom",
              "tokenizer" : "keyword"
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "W2Gu9s9JQMWKgANfAtZGbA",
        "version" : {
          "created" : "5010199"
        }
      }
    }
  }
}`)
)

type FakeawsESIndexReaper struct {
	indexList indexes
}

func (es FakeawsESIndexReaper) esRequest(url, action string) ([]byte, error) {
	return fakeResults, nil
}

func TestSortIndexList(t *testing.T) {
	expectedCount := 28
	maxCount = 5
	f := FakeawsESIndexReaper{}

	body, _ := makeESRequest(f, "url", "GET")

	err := json.Unmarshal(body, &indexReaper.indexList)
	if err != nil {
		t.Fatalf("Failed to unmarshal index list into struct: %s", err)
	}

	sortedIndexes, err := indexReaper.sortIndexesByAge()

	if expectedCount != len(sortedIndexes) {
		t.Fatalf("Expected index count of %d but got %d", expectedCount, len(sortedIndexes))
	}

	if sortedIndexes[0].Settings.Index.ProvidedName != "zipkin-2017-02-14" {
		t.Fatalf("Expected first index to be %s but got %s", "zipkin-2017-02-14", sortedIndexes[0].Settings.Index.ProvidedName)
	}

	if sortedIndexes[27].Settings.Index.ProvidedName != "zipkin-2017-03-13" {
		t.Fatalf("Expected last index to be %s but got %s", "zipkin-2017-03-13", sortedIndexes[27].Settings.Index.ProvidedName)
	}
}
