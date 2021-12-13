Type point
    x As Integer
    y As Integer
End Type
Open "c:\tmp\day13.txt" For Input As #1
linenr! = 0
Dim lines(750) As point
While Not EOF(1)
    Line Input #1, line$
    If InStr(line$, ",") > 0 Then
        Dim p As point
        p.x = Val(Left$(line$, InStr(line$, ",") - 1))
        p.y = Val(Right$(line$, Len(line$) - InStr(line$, ",")))
        lines(linenr!) = p
        linenr! = linenr! + 1
    End If
    If InStr(line$, "=") > 0 Then
        fold! = Val(Right$(line$, Len(line$) - InStr(line$, "=")))
        For i = 0 To linenr! - 1
            p = lines(i)
            If Mid$(line$, 12, 1) = "x" And p.x > fold! Then
                p.x = fold! - (p.x - fold!)
            End If
            If Mid$(line$, 12, 1) = "y" And p.y > fold! Then
                p.y = fold! - (p.y - fold!)
            End If
            lines(i) = p
        Next
    End If
Wend
Close #1
For y = 0 To 7
    line$ = ""
    For x = 0 To 40
        found! = 0
        For i = 0 To linenr! - 1
            p = lines(i)
            If p.x = x And p.y = y Then
                found! = 1
            End If
        Next i
        If found! = 1 Then
            line$ = line$ + "#"
        Else
            line$ = line$ + " "
        End If
    Next x
    Print line$
Next y






