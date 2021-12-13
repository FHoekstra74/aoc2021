Open "c:\tmp\day1.txt" For Input As #1
a! = 0
b! = 0
linenr! = 0
Dim lines(2000) As Integer
While Not EOF(1)
    Input #1, line$
    number! = Val(line$)
    lines(linenr!) = number!
    linenr! = linenr! + 1
Wend
Close #1
For i = 1 To 2000
    If lines(i) > lines(i - 1) Then
        a! = a! + 1
    End If
    If i > 2 Then
        If lines(i) > lines(i - 3) Then
            b! = b! + 1
        End If
    End If
    Print lines(i)
Next
Print "Answer a:"; a!
Print "Answer b:"; b!




