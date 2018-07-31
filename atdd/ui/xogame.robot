***Settings***
Library    SeleniumLibrary
Suite Teardown    ปิด Browser

***Variables***
${URL}    http://localhost:3000/xogame

***Test Cases***
เล่นเกม xo ระหว่างกาดกับเลขในตาที่ 5 กาดชนะแนวนอน
    เปิด Browser
    กรอกชื่อผู้เล่นคนแรก    กาด
    กรอกชื่อผู้เล่นคนที่สอง    เล็ก
    กดปุ่ม OK
    เข้าหน้าเล่นเกม
    Turn ของ    กาด
    เลือกช่อง    1    2    X
    Turn ของ    เล็ก
    เลือกช่อง    2    2    O
    Turn ของ    กาด
    เลือกช่อง    1    1    X
    Turn ของ    เล็ก
    เลือกช่อง    2    1    O
    Turn ของ    กาด
    เลือกช่อง    1    3    X
    แสดงผู้ชนะ    กาด
    กลับไปยังหน้าแรก

***Keywords***
เปิด Browser
    Open Browser    ${URL}    chrome
กรอกชื่อผู้เล่นคนแรก
    [Arguments]    ${NAME}
    Input Text    id=playerX    ${NAME}
กรอกชื่อผู้เล่นคนที่สอง
    [Arguments]    ${NAME}
    Input Text    id=playerO    ${NAME}
กดปุ่ม OK
    Click Button    id=startGame
เข้าหน้าเล่นเกม
    Element Text Should Be    id=playerX    กาด
Turn ของ
    [Arguments]    ${NAME}
    Element Text Should Be    id=turnOf    ${NAME}
เลือกช่อง
    [Arguments]    ${ROW}    ${COLUMN}    ${SYMBOL}
    Click Element    id=grid-${ROW}-${COLUMN}
    Wait Until Element Contains    id=grid-${ROW}-${COLUMN}    ${SYMBOL}
แสดงผู้ชนะ
    [Arguments]    ${NAME}
    Wait Until Element Contains    id=winner    ${NAME}
กลับไปยังหน้าแรก
    Click Button    id=ok
ปิด Browser
    Close Browser