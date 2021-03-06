// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef SKYWALKING_SPAN_H
#define SKYWALKING_SPAN_H

#include <string>
#include <vector>
#include "tag.h"
#include "segment_reference.h"

enum class SkySpanType {
    Entry, Exit, Local
};

enum class SkySpanLayer {
    Unknown,
    Database,
    RPCFramework,
    Http,
    MQ,
    Cache
};

class Span {
public:
    Span();
    
    ~Span();

    int getSpanId() const;

    int getParentSpanId() const;

    uint64_t getStartTime() const;

    long getEndTime() const;

    std::vector<SkySegmentReference *> getRefs();

    std::string getOperationName();

    std::string getPeer();

    SkySpanType getSpanType();

    SkySpanLayer getSpanLayer();

    int getComponentId() const;

    bool getIsError() const;

    std::vector<Tag *> getTags();

    bool getSkipAnalysis() const;

    // //
    void setEndTIme();

    void setOperationName(const std::string &name);

    void setPeer(const std::string &peer);

    void setSpanType(SkySpanType type);

    void setSpanLayer(SkySpanLayer layer);

    void setComponentId(int id);

    void setIsError(bool error);

    void setSpanId(int id);

    void setParentSpanId(int id);

    void pushTag(Tag *tag);

    void addTag(const std::string &key, const std::string &value);

    void pushRefs(SkySegmentReference *ref);

private:
    static uint64_t getUnixTimeStamp();

private:
    int spanId;
    int parentSpanId;
    uint64_t startTime;
    long endTime;
    std::vector<SkySegmentReference *> refs;
    std::string operationName;
    std::string peer;
    SkySpanType spanType;
    SkySpanLayer spanLayer;
    int componentId;
    bool isError;
    std::vector<Tag *> tags;
    bool skipAnalysis;
};


#endif //SKYWALKING_SPAN_H
